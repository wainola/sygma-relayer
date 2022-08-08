// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/calls/events/listener.go

// Package mock_events is a generated GoMock package.
package mock_events

import (
	context "context"
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockChainClient is a mock of ChainClient interface.
type MockChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockChainClientMockRecorder
}

// MockChainClientMockRecorder is the mock recorder for MockChainClient.
type MockChainClientMockRecorder struct {
	mock *MockChainClient
}

// NewMockChainClient creates a new mock instance.
func NewMockChainClient(ctrl *gomock.Controller) *MockChainClient {
	mock := &MockChainClient{ctrl: ctrl}
	mock.recorder = &MockChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainClient) EXPECT() *MockChainClientMockRecorder {
	return m.recorder
}

// FetchEventLogs mocks base method.
func (m *MockChainClient) FetchEventLogs(ctx context.Context, contractAddress common.Address, event string, startBlock, endBlock *big.Int) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEventLogs", ctx, contractAddress, event, startBlock, endBlock)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEventLogs indicates an expected call of FetchEventLogs.
func (mr *MockChainClientMockRecorder) FetchEventLogs(ctx, contractAddress, event, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEventLogs", reflect.TypeOf((*MockChainClient)(nil).FetchEventLogs), ctx, contractAddress, event, startBlock, endBlock)
}

// LatestBlock mocks base method.
func (m *MockChainClient) LatestBlock() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlock")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlock indicates an expected call of LatestBlock.
func (mr *MockChainClientMockRecorder) LatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlock", reflect.TypeOf((*MockChainClient)(nil).LatestBlock))
}

// TransactionByHash mocks base method.
func (m *MockChainClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockChainClientMockRecorder) TransactionByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockChainClient)(nil).TransactionByHash), ctx, hash)
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockChainClient) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockChainClientMockRecorder) WaitAndReturnTxReceipt(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockChainClient)(nil).WaitAndReturnTxReceipt), h)
}
