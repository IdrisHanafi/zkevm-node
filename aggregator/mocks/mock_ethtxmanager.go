// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	coretypes "github.com/ethereum/go-ethereum/core/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/etherman/types"
)

// EthTxManager is an autogenerated mock type for the ethTxManager type
type EthTxManager struct {
	mock.Mock
}

// VerifyBatches provides a mock function with given fields: ctx, lastVerifiedBatch, batchNum, inputs
func (_m *EthTxManager) VerifyBatches(ctx context.Context, lastVerifiedBatch uint64, batchNum uint64, inputs *types.FinalProofInputs) (*coretypes.Transaction, error) {
	ret := _m.Called(ctx, lastVerifiedBatch, batchNum, inputs)

	var r0 *coretypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, *types.FinalProofInputs) *coretypes.Transaction); ok {
		r0 = rf(ctx, lastVerifiedBatch, batchNum, inputs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, *types.FinalProofInputs) error); ok {
		r1 = rf(ctx, lastVerifiedBatch, batchNum, inputs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEthTxManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewEthTxManager creates a new instance of EthTxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEthTxManager(t mockConstructorTestingTNewEthTxManager) *EthTxManager {
	mock := &EthTxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
