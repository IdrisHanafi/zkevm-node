package main

import (
	"encoding/binary"
	"os"
	"reflect"
	"time"

	"github.com/0xPolygonHermez/zkevm-data-streamer/datastreamer"
	"github.com/0xPolygonHermez/zkevm-data-streamer/log"
	"github.com/0xPolygonHermez/zkevm-node/db"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/0xPolygonHermez/zkevm-node/state/runtime/executor"
	"github.com/0xPolygonHermez/zkevm-node/tools/datastreamer/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

const appName = "zkevm-data-streamer-tool"

var (
	configFileFlag = cli.StringFlag{
		Name:     config.FlagCfg,
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: false,
	}

	entryFlag = cli.Uint64Flag{
		Name:     "entry",
		Aliases:  []string{"e"},
		Usage:    "Entry `NUMBER`",
		Required: false,
	}

	l2blockFlag = cli.Uint64Flag{
		Name:     "l2block",
		Aliases:  []string{"b"},
		Usage:    "L2Block `NUMBER`",
		Required: false,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = appName

	flags := []cli.Flag{
		&configFileFlag,
		&entryFlag,
		&l2blockFlag,
	}

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{},
			Usage:   "Generate stream file form scratch",
			Action:  generate,
			Flags:   flags,
		},
		{
			Name:    "rebuild",
			Aliases: []string{},
			Usage:   "Rebuild state roots from a block",
			Action:  rebuild,
			Flags:   flags,
		},
		{
			Name:    "decode-entry",
			Aliases: []string{},
			Usage:   "Decodes an entry",
			Action:  decodeEntry,
			Flags:   flags,
		},
		{
			Name:    "decode-l2block",
			Aliases: []string{},
			Usage:   "Decodes a l2 block",
			Action:  decodeL2Block,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initializeStreamServer(c *config.Config) (*datastreamer.StreamServer, error) {
	// Init logger
	log.Init(c.StreamServer.Log)
	log.Info("Starting tool")

	// Create a stream server
	streamServer, err := datastreamer.New(c.StreamServer.Port, state.StreamTypeSequencer, c.StreamServer.Filename, &c.StreamServer.Log)
	if err != nil {
		return nil, err
	}

	// Set entities definition
	entriesDefinition := map[datastreamer.EntryType]datastreamer.EntityDefinition{
		state.EntryTypeL2BlockStart: {
			Name:       "L2BlockStart",
			StreamType: state.StreamTypeSequencer,
			Definition: reflect.TypeOf(state.DSL2BlockStart{}),
		},
		state.EntryTypeL2Tx: {
			Name:       "L2Transaction",
			StreamType: state.StreamTypeSequencer,
			Definition: reflect.TypeOf(state.DSL2Transaction{}),
		},
		state.EntryTypeL2BlockEnd: {
			Name:       "L2BlockEnd",
			StreamType: state.StreamTypeSequencer,
			Definition: reflect.TypeOf(state.DSL2BlockEnd{}),
		},
	}

	streamServer.SetEntriesDef(entriesDefinition)

	err = streamServer.Start()
	if err != nil {
		return nil, err
	}

	return &streamServer, nil
}

func generate(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Loaded configuration: %+v", c)

	streamServer, err := initializeStreamServer(c)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	stateSqlDB, err := db.NewSQLDB(c.StateDB)
	if err != nil {
		log.Fatal(err)
	}
	defer stateSqlDB.Close()
	stateDB := state.NewPostgresStorage(stateSqlDB)
	log.Info("Connected to the database")

	header := streamServer.GetHeader()

	var currentL2Block uint64
	var currentTxIndex uint64

	if header.TotalEntries == 0 {
		// Get Genesis block
		genesisL2Block, err := stateDB.GetDSGenesisBlock(cliCtx.Context, nil)
		if err != nil {
			log.Fatal(err)
		}

		err = streamServer.StartAtomicOp()
		if err != nil {
			log.Fatal(err)
		}

		bookMark := state.DSBookMark{
			Type:          state.BookMarkTypeL2Block,
			L2BlockNumber: genesisL2Block.L2BlockNumber,
		}

		_, err = streamServer.AddStreamBookmark(bookMark.Encode())
		if err != nil {
			log.Fatal(err)
		}

		genesisBlock := state.DSL2BlockStart{
			BatchNumber:    genesisL2Block.BatchNumber,
			L2BlockNumber:  genesisL2Block.L2BlockNumber,
			Timestamp:      genesisL2Block.Timestamp,
			GlobalExitRoot: genesisL2Block.GlobalExitRoot,
			Coinbase:       genesisL2Block.Coinbase,
			ForkID:         genesisL2Block.ForkID,
		}

		log.Infof("Genesis block: %+v", genesisBlock)

		_, err = streamServer.AddStreamEntry(1, genesisBlock.Encode())
		if err != nil {
			log.Fatal(err)
		}

		genesisBlockEnd := state.DSL2BlockEnd{
			L2BlockNumber: genesisL2Block.L2BlockNumber,
			BlockHash:     genesisL2Block.BlockHash,
			StateRoot:     genesisL2Block.StateRoot,
		}

		_, err = streamServer.AddStreamEntry(state.EntryTypeL2BlockEnd, genesisBlockEnd.Encode())
		if err != nil {
			log.Fatal(err)
		}

		err = streamServer.CommitAtomicOp()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		latestEntry, err := streamServer.GetEntry(header.TotalEntries - 1)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Latest entry: %+v", latestEntry)

		switch latestEntry.Type {
		case state.EntryTypeL2BlockStart:
			log.Info("Latest entry type is L2BlockStart")
			currentL2Block = binary.LittleEndian.Uint64(latestEntry.Data[8:16])
		case state.EntryTypeL2Tx:
			log.Info("Latest entry type is L2Tx")

			for latestEntry.Type == state.EntryTypeL2Tx {
				currentTxIndex++
				latestEntry, err = streamServer.GetEntry(header.TotalEntries - currentTxIndex)
				if err != nil {
					log.Fatal(err)
				}
			}

			if latestEntry.Type != state.EntryTypeL2BlockStart {
				log.Fatal("Latest entry is not a L2BlockStart")
			}
			currentL2Block = binary.LittleEndian.Uint64(latestEntry.Data[8:16])

		case state.EntryTypeL2BlockEnd:
			log.Info("Latest entry type is L2BlockEnd")
			currentL2Block = binary.LittleEndian.Uint64(latestEntry.Data[0:8])
		}
	}

	log.Infof("Current transaction index: %d", currentTxIndex)
	log.Infof("Current L2 block number: %d", currentL2Block)

	var limit uint64 = 5000
	var offset uint64 = currentL2Block
	var entry uint64 = header.TotalEntries
	var l2blocks []*state.DSL2Block

	if entry > 0 {
		entry--
	}

	for err == nil {
		log.Infof("Current entry number: %d", entry)

		l2blocks, err = stateDB.GetDSL2Blocks(cliCtx.Context, limit, offset, nil)
		offset += limit
		if len(l2blocks) == 0 {
			break
		}
		// Get transactions for all the retrieved l2 blocks
		l2Transactions, err := stateDB.GetDSL2Transactions(cliCtx.Context, l2blocks[0].L2BlockNumber, l2blocks[len(l2blocks)-1].L2BlockNumber, nil)
		if err != nil {
			log.Fatal(err)
		}

		err = streamServer.StartAtomicOp()
		if err != nil {
			log.Fatal(err)
		}

		for x, l2block := range l2blocks {
			if currentTxIndex > 0 {
				x += int(currentTxIndex)
				currentTxIndex = 0
			}

			blockStart := state.DSL2BlockStart{
				BatchNumber:    l2block.BatchNumber,
				L2BlockNumber:  l2block.L2BlockNumber,
				Timestamp:      l2block.Timestamp,
				GlobalExitRoot: l2block.GlobalExitRoot,
				Coinbase:       l2block.Coinbase,
				ForkID:         l2block.ForkID,
			}

			bookMark := state.DSBookMark{
				Type:          state.BookMarkTypeL2Block,
				L2BlockNumber: blockStart.L2BlockNumber,
			}

			_, err = streamServer.AddStreamBookmark(bookMark.Encode())
			if err != nil {
				log.Fatal(err)
			}

			_, err = streamServer.AddStreamEntry(state.EntryTypeL2BlockStart, blockStart.Encode())
			if err != nil {
				log.Fatal(err)
			}

			entry, err = streamServer.AddStreamEntry(state.EntryTypeL2Tx, l2Transactions[x].Encode())
			if err != nil {
				log.Fatal(err)
			}

			blockEnd := state.DSL2BlockEnd{
				L2BlockNumber: l2block.L2BlockNumber,
				BlockHash:     l2block.BlockHash,
				StateRoot:     l2block.StateRoot,
			}

			_, err = streamServer.AddStreamEntry(state.EntryTypeL2BlockEnd, blockEnd.Encode())
			if err != nil {
				log.Fatal(err)
			}
		}
		err = streamServer.CommitAtomicOp()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Finished tool")

	return nil
}

func rebuild(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Loaded configuration: %+v", c)

	ctx := cliCtx.Context

	streamServer, err := initializeStreamServer(c)
	if err != nil {
		log.Fatal(err)
	}

	oldEndEntry, err := streamServer.GetEntry(1)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("oldEndEntry: %+v", oldEndEntry)

	startEntry, err := streamServer.GetEntry(2) //nolint:gomnd
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("startEntry: %+v", startEntry)
	log.Infof("Length of data in startEntry: %d", len(startEntry.Data))

	txEntry, err := streamServer.GetEntry(3) //nolint:gomnd
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("txEntry: %+v", txEntry)

	endEntry, err := streamServer.GetEntry(4) //nolint:gomnd
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("endEntry: %+v", endEntry)

	tx, err := state.DecodeTx(string(txEntry.Data[6:]))
	if err != nil {
		log.Fatal(err)
	}

	/*
		log.Infof("tx nonce: %+v", tx.Nonce())

		sender, err := state.GetSender(*tx)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("tx sender: %+v", sender)
	*/

	// RLP encode the transaction using the proper fork id
	batchL2Data, err := state.EncodeTransaction(*tx, 255, uint64(binary.LittleEndian.Uint16(startEntry.Data[76:78]))) //nolint:gomnd
	if err != nil {
		log.Fatal(err)
	}

	processBatchRequest := &executor.ProcessBatchRequest{
		OldBatchNum:      uint64(startEntry.Data[0]) - 1,
		Coinbase:         common.Bytes2Hex(startEntry.Data[56:76]),
		BatchL2Data:      batchL2Data,
		OldStateRoot:     oldEndEntry.Data[40:72],
		GlobalExitRoot:   startEntry.Data[23:55],
		OldAccInputHash:  []byte{},
		EthTimestamp:     binary.LittleEndian.Uint64(startEntry.Data[16:24]),
		UpdateMerkleTree: uint32(0),
		ChainId:          1001, //nolint:gomnd
		ForkId:           uint64(binary.LittleEndian.Uint16(startEntry.Data[76:78])),
	}

	log.Infof("ProcessBatchRequest: %+v", processBatchRequest)

	log.Infof("Old state root:    %s", common.Bytes2Hex(processBatchRequest.OldStateRoot))
	log.Infof("Expected new root: %s", common.Bytes2Hex(endEntry.Data[40:72]))
	log.Infof("Batch L2 data:     %s", common.Bytes2Hex(processBatchRequest.BatchL2Data))
	log.Infof("Coinbase:          %s", processBatchRequest.Coinbase)
	log.Infof("Timestamp:         %d", processBatchRequest.EthTimestamp)
	log.Infof("Fork id:           %d", processBatchRequest.ForkId)

	/*
		if processBatchRequest.ForkId >= 5 {
			processBatchRequest.BatchL2Data = append(processBatchRequest.BatchL2Data, byte(255))
		}
	*/

	// Connect to the executor
	executorClient, executorClientConn, executorCancel := executor.NewExecutorClient(ctx, c.Executor)
	s := executorClientConn.GetState()
	log.Infof("executorClientConn state: %s", s.String())
	defer func() {
		executorCancel()
		executorClientConn.Close()
	}()

	/*
		mtDBServerConfig := merkletree.Config{URI: fmt.Sprintf("%s:50061", "toni-prover")}
		var mtDBCancel context.CancelFunc
		mtDBServiceClient, mtDBClientConn, mtDBCancel := merkletree.NewMTDBServiceClient(ctx, mtDBServerConfig)
		s = mtDBClientConn.GetState()
		log.Infof("stateDbClientConn state: %s", s.String())
		defer func() {
			mtDBCancel()
			mtDBClientConn.Close()
		}()

		stateTree := merkletree.NewStateTree(mtDBServiceClient)

		// Connect to MT
		nonce, err := stateTree.GetNonce(ctx, sender, processBatchRequest.OldStateRoot)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Nonce in MT: %d", nonce)
	*/

	// Process batch
	processBatchResponse, err := executorClient.ProcessBatch(ctx, processBatchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if processBatchResponse.Error != executor.ExecutorError_EXECUTOR_ERROR_NO_ERROR {
		log.Fatal(processBatchResponse.Error)
	}

	log.Infof("ProcessBatchResponse: %+v", processBatchResponse)

	log.Infof("New root: %s", common.Bytes2Hex(processBatchResponse.NewStateRoot))

	return nil
}

func decodeEntry(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Loaded configuration: %+v", c)

	streamServer, err := initializeStreamServer(c)
	if err != nil {
		log.Fatal(err)
	}

	entry, err := streamServer.GetEntry(cliCtx.Uint64("entry"))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Selected entry: %+v", entry)

	printEntry(entry)

	return nil
}

func decodeL2Block(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Loaded configuration: %+v", c)

	streamServer, err := initializeStreamServer(c)
	if err != nil {
		log.Fatal(err)
	}

	l2BlockNumber := cliCtx.Uint64("l2block")

	bookMark := state.DSBookMark{
		Type:          state.BookMarkTypeL2Block,
		L2BlockNumber: l2BlockNumber,
	}

	firstEntry, err := streamServer.GetFirstEventAfterBookmark(bookMark.Encode())
	if err != nil {
		log.Fatal(err)
	}
	printEntry(firstEntry)

	secondEntry, err := streamServer.GetEntry(firstEntry.Number + 1)
	if err != nil {
		log.Fatal(err)
	}
	printEntry(secondEntry)

	if l2BlockNumber != 0 {
		thirdEntry, err := streamServer.GetEntry(firstEntry.Number + 2) //nolint:gomnd
		if err != nil {
			log.Fatal(err)
		}
		printEntry(thirdEntry)
	}

	return nil
}

func printEntry(entry datastreamer.FileEntry) {
	switch entry.Type {
	case state.EntryTypeBookMark:
		log.Infof("Entry %d: BookMark", entry.Number)
		l2BlockNumber := binary.LittleEndian.Uint64(entry.Data[1:9])
		log.Infof("L2 block number: %d", l2BlockNumber)
	case state.EntryTypeL2BlockStart:
		log.Infof("Entry %d: L2BlockStart", entry.Number)
		batchNumber := binary.LittleEndian.Uint64(entry.Data[0:8])
		log.Infof("Batch number: %d", batchNumber)
		l2BlockNumber := binary.LittleEndian.Uint64(entry.Data[8:16])
		log.Infof("L2 block number: %d", l2BlockNumber)
		timestamp := binary.LittleEndian.Uint64(entry.Data[16:24])
		log.Infof("Timestamp: %d", timestamp)
		log.Infof("Timestamp: %v", time.Unix(int64(timestamp), 0))
		globalExitRoot := "0x" + common.Bytes2Hex(entry.Data[24:56])
		log.Infof("Global exit root: %s", globalExitRoot)
		coinbase := "0x" + common.Bytes2Hex(entry.Data[56:76])
		log.Infof("Coinbase: %s", coinbase)
		forkID := binary.LittleEndian.Uint16(entry.Data[76:78])
		log.Infof("Fork ID: %d", forkID)
	case state.EntryTypeL2Tx:
		log.Infof("Entry %d: L2Tx", entry.Number)
		effectiveGasPricePercentage := entry.Data[0]
		log.Infof("Effective gas price percentage: %d", effectiveGasPricePercentage)
		isValid := entry.Data[1] == 1
		log.Infof("Is valid: %t", isValid)
		encodeLength := binary.LittleEndian.Uint16(entry.Data[2:6])
		log.Infof("Encode length: %d", encodeLength)
		encode := entry.Data[6:]
		log.Infof("Encode: %s", "0x"+common.Bytes2Hex(encode))
		tx, err := state.DecodeTx(common.Bytes2Hex(encode))
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Decoded: %+v", tx)
	case state.EntryTypeL2BlockEnd:
		log.Infof("Entry %d: L2BlockEnd", entry.Number)
		log.Infof("L2 Block Number: %d", binary.LittleEndian.Uint64(entry.Data[0:8]))
		log.Infof("Block Hash: %s", "0x"+common.Bytes2Hex(entry.Data[8:40]))
		log.Infof("State root: %s", "0x"+common.Bytes2Hex(entry.Data[40:72]))
	}
}
