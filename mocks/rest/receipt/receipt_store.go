// Code generated by mockery v2.29.0. DO NOT EDIT.

package mockreceipt

import (
	http "net/http"

	api "github.com/hyperledger/firefly-fabconnect/internal/rest/receipt/api"

	httprouter "github.com/julienschmidt/httprouter"

	mock "github.com/stretchr/testify/mock"

	ws "github.com/hyperledger/firefly-fabconnect/internal/ws"
)

// ReceiptStore is an autogenerated mock type for the ReceiptStore type
type ReceiptStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ReceiptStore) Close() {
	_m.Called()
}

// GetReceipt provides a mock function with given fields: res, req, params
func (_m *ReceiptStore) GetReceipt(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_m.Called(res, req, params)
}

// GetReceipts provides a mock function with given fields: res, req, params
func (_m *ReceiptStore) GetReceipts(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	_m.Called(res, req, params)
}

// Init provides a mock function with given fields: _a0, _a1
func (_m *ReceiptStore) Init(_a0 ws.WebSocketChannels, _a1 ...api.ReceiptStorePersistence) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(ws.WebSocketChannels, ...api.ReceiptStorePersistence) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessReceipt provides a mock function with given fields: msgBytes
func (_m *ReceiptStore) ProcessReceipt(msgBytes []byte) {
	_m.Called(msgBytes)
}

// ValidateConf provides a mock function with given fields:
func (_m *ReceiptStore) ValidateConf() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewReceiptStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewReceiptStore creates a new instance of ReceiptStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReceiptStore(t mockConstructorTestingTNewReceiptStore) *ReceiptStore {
	mock := &ReceiptStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
