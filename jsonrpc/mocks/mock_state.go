// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	runtime "github.com/0xPolygonHermez/zkevm-node/state/runtime"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"
)

// StateMock is an autogenerated mock type for the StateInterface type
type StateMock struct {
	mock.Mock
}

// BatchNumberByL2BlockNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) BatchNumberByL2BlockNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DebugTransaction provides a mock function with given fields: ctx, transactionHash, traceConfig, dbTx
func (_m *StateMock) DebugTransaction(ctx context.Context, transactionHash common.Hash, traceConfig state.TraceConfig, dbTx pgx.Tx) (*runtime.ExecutionResult, error) {
	ret := _m.Called(ctx, transactionHash, traceConfig, dbTx)

	var r0 *runtime.ExecutionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, state.TraceConfig, pgx.Tx) (*runtime.ExecutionResult, error)); ok {
		return rf(ctx, transactionHash, traceConfig, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, state.TraceConfig, pgx.Tx) *runtime.ExecutionResult); ok {
		r0 = rf(ctx, transactionHash, traceConfig, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ExecutionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, state.TraceConfig, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, traceConfig, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: transaction, senderAddress, l2BlockNumber, dbTx
func (_m *StateMock) EstimateGas(transaction *coretypes.Transaction, senderAddress common.Address, l2BlockNumber *uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(transaction, senderAddress, l2BlockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(*coretypes.Transaction, common.Address, *uint64, pgx.Tx) (uint64, error)); ok {
		return rf(transaction, senderAddress, l2BlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(*coretypes.Transaction, common.Address, *uint64, pgx.Tx) uint64); ok {
		r0 = rf(transaction, senderAddress, l2BlockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(*coretypes.Transaction, common.Address, *uint64, pgx.Tx) error); ok {
		r1 = rf(transaction, senderAddress, l2BlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, address, l2Block
func (_m *StateMock) GetBalance(ctx context.Context, address common.Address, l2Block *coretypes.Block) (*big.Int, error) {
	ret := _m.Called(ctx, address, l2Block)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *coretypes.Block) (*big.Int, error)); ok {
		return rf(ctx, address, l2Block)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *coretypes.Block) *big.Int); ok {
		r0 = rf(ctx, address, l2Block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *coretypes.Block) error); ok {
		r1 = rf(ctx, address, l2Block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCode provides a mock function with given fields: ctx, address, l2Block
func (_m *StateMock) GetCode(ctx context.Context, address common.Address, l2Block *coretypes.Block) ([]byte, error) {
	ret := _m.Called(ctx, address, l2Block)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *coretypes.Block) ([]byte, error)); ok {
		return rf(ctx, address, l2Block)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *coretypes.Block) []byte); ok {
		r0 = rf(ctx, address, l2Block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *coretypes.Block) error); ok {
		r1 = rf(ctx, address, l2Block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExitRootByGlobalExitRoot provides a mock function with given fields: ctx, ger, dbTx
func (_m *StateMock) GetExitRootByGlobalExitRoot(ctx context.Context, ger common.Hash, dbTx pgx.Tx) (*state.GlobalExitRoot, error) {
	ret := _m.Called(ctx, ger, dbTx)

	var r0 *state.GlobalExitRoot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (*state.GlobalExitRoot, error)); ok {
		return rf(ctx, ger, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *state.GlobalExitRoot); ok {
		r0 = rf(ctx, ger, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.GlobalExitRoot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, ger, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockByHash provides a mock function with given fields: ctx, hash, dbTx
func (_m *StateMock) GetL2BlockByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (*coretypes.Block, error) {
	ret := _m.Called(ctx, hash, dbTx)

	var r0 *coretypes.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (*coretypes.Block, error)); ok {
		return rf(ctx, hash, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *coretypes.Block); ok {
		r0 = rf(ctx, hash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, hash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) GetL2BlockByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*coretypes.Block, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 *coretypes.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*coretypes.Block, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *coretypes.Block); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockHashesSince provides a mock function with given fields: ctx, since, dbTx
func (_m *StateMock) GetL2BlockHashesSince(ctx context.Context, since time.Time, dbTx pgx.Tx) ([]common.Hash, error) {
	ret := _m.Called(ctx, since, dbTx)

	var r0 []common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, pgx.Tx) ([]common.Hash, error)); ok {
		return rf(ctx, since, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, pgx.Tx) []common.Hash); ok {
		r0 = rf(ctx, since, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, pgx.Tx) error); ok {
		r1 = rf(ctx, since, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockHeaderByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) GetL2BlockHeaderByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*coretypes.Header, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 *coretypes.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*coretypes.Header, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *coretypes.Header); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockTransactionCountByHash provides a mock function with given fields: ctx, hash, dbTx
func (_m *StateMock) GetL2BlockTransactionCountByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, hash, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, hash, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) uint64); ok {
		r0 = rf(ctx, hash, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, hash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockTransactionCountByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) GetL2BlockTransactionCountByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastConsolidatedL2BlockNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastConsolidatedL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2Block provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastL2Block(ctx context.Context, dbTx pgx.Tx) (*coretypes.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *coretypes.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*coretypes.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *coretypes.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2BlockNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVerifiedBatch provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVerifiedBatch(ctx context.Context, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.VerifiedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.VerifiedBatch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx
func (_m *StateMock) GetLogs(ctx context.Context, fromBlock uint64, toBlock uint64, addresses []common.Address, topics [][]common.Hash, blockHash *common.Hash, since *time.Time, dbTx pgx.Tx) ([]*coretypes.Log, error) {
	ret := _m.Called(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)

	var r0 []*coretypes.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, pgx.Tx) ([]*coretypes.Log, error)); ok {
		return rf(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, pgx.Tx) []*coretypes.Log); ok {
		r0 = rf(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*coretypes.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, pgx.Tx) error); ok {
		r1 = rf(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address, blockNumber, dbTx
func (_m *StateMock) GetNonce(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, address, blockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, address, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageAt provides a mock function with given fields: ctx, address, position, l2Block
func (_m *StateMock) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, l2Block *coretypes.Block) (*big.Int, error) {
	ret := _m.Called(ctx, address, position, l2Block)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, *coretypes.Block) (*big.Int, error)); ok {
		return rf(ctx, address, position, l2Block)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, *coretypes.Block) *big.Int); ok {
		r0 = rf(ctx, address, position, l2Block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int, *coretypes.Block) error); ok {
		r1 = rf(ctx, address, position, l2Block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncingInfo provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetSyncingInfo(ctx context.Context, dbTx pgx.Tx) (state.SyncingInfo, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 state.SyncingInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (state.SyncingInfo, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) state.SyncingInfo); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(state.SyncingInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByHash provides a mock function with given fields: ctx, transactionHash, dbTx
func (_m *StateMock) GetTransactionByHash(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*coretypes.Transaction, error) {
	ret := _m.Called(ctx, transactionHash, dbTx)

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (*coretypes.Transaction, error)); ok {
		return rf(ctx, transactionHash, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *coretypes.Transaction); ok {
		r0 = rf(ctx, transactionHash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByL2BlockHashAndIndex provides a mock function with given fields: ctx, blockHash, index, dbTx
func (_m *StateMock) GetTransactionByL2BlockHashAndIndex(ctx context.Context, blockHash common.Hash, index uint64, dbTx pgx.Tx) (*coretypes.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index, dbTx)

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint64, pgx.Tx) (*coretypes.Transaction, error)); ok {
		return rf(ctx, blockHash, index, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint64, pgx.Tx) *coretypes.Transaction); ok {
		r0 = rf(ctx, blockHash, index, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockHash, index, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByL2BlockNumberAndIndex provides a mock function with given fields: ctx, blockNumber, index, dbTx
func (_m *StateMock) GetTransactionByL2BlockNumberAndIndex(ctx context.Context, blockNumber uint64, index uint64, dbTx pgx.Tx) (*coretypes.Transaction, error) {
	ret := _m.Called(ctx, blockNumber, index, dbTx)

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) (*coretypes.Transaction, error)); ok {
		return rf(ctx, blockNumber, index, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) *coretypes.Transaction); ok {
		r0 = rf(ctx, blockNumber, index, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, index, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionReceipt provides a mock function with given fields: ctx, transactionHash, dbTx
func (_m *StateMock) GetTransactionReceipt(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*coretypes.Receipt, error) {
	ret := _m.Called(ctx, transactionHash, dbTx)

	var r0 *coretypes.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) (*coretypes.Receipt, error)); ok {
		return rf(ctx, transactionHash, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *coretypes.Receipt); ok {
		r0 = rf(ctx, transactionHash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]coretypes.Transaction, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) ([]coretypes.Transaction, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []coretypes.Transaction); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVerifiedBatch provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetVerifiedBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.VerifiedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.VerifiedBatch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVirtualBatch provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetVirtualBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VirtualBatch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.VirtualBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.VirtualBatch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.VirtualBatch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VirtualBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsL2BlockConsolidated provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) IsL2BlockConsolidated(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (bool, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsL2BlockVirtualized provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) IsL2BlockVirtualized(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (bool, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareWebSocket provides a mock function with given fields:
func (_m *StateMock) PrepareWebSocket() {
	_m.Called()
}

// ProcessUnsignedTransaction provides a mock function with given fields: ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx
func (_m *StateMock) ProcessUnsignedTransaction(ctx context.Context, tx *coretypes.Transaction, senderAddress common.Address, l2BlockNumber *uint64, noZKEVMCounters bool, dbTx pgx.Tx) (*runtime.ExecutionResult, error) {
	ret := _m.Called(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)

	var r0 *runtime.ExecutionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Transaction, common.Address, *uint64, bool, pgx.Tx) (*runtime.ExecutionResult, error)); ok {
		return rf(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Transaction, common.Address, *uint64, bool, pgx.Tx) *runtime.ExecutionResult); ok {
		r0 = rf(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ExecutionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *coretypes.Transaction, common.Address, *uint64, bool, pgx.Tx) error); ok {
		r1 = rf(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNewL2BlockEventHandler provides a mock function with given fields: h
func (_m *StateMock) RegisterNewL2BlockEventHandler(h state.NewL2BlockEventHandler) {
	_m.Called(h)
}

type mockConstructorTestingTNewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateMock creates a new instance of StateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateMock(t mockConstructorTestingTNewStateMock) *StateMock {
	mock := &StateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
