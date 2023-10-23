// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Do not include this in mocks.mockgen.txt as bls package won't be available.
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/validators (interfaces: Manager)

// Package validators is a generated GoMock package.
package validators

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	bls "github.com/ava-labs/avalanchego/utils/crypto/bls"
	set "github.com/ava-labs/avalanchego/utils/set"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddStaker mocks base method.
func (m *MockManager) AddStaker(arg0 ids.ID, arg1 ids.NodeID, arg2 *bls.PublicKey, arg3 ids.ID, arg4 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStaker", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStaker indicates an expected call of AddStaker.
func (mr *MockManagerMockRecorder) AddStaker(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStaker", reflect.TypeOf((*MockManager)(nil).AddStaker), arg0, arg1, arg2, arg3, arg4)
}

// AddWeight mocks base method.
func (m *MockManager) AddWeight(arg0 ids.ID, arg1 ids.NodeID, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWeight", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWeight indicates an expected call of AddWeight.
func (mr *MockManagerMockRecorder) AddWeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWeight", reflect.TypeOf((*MockManager)(nil).AddWeight), arg0, arg1, arg2)
}

// Contains mocks base method.
func (m *MockManager) Contains(arg0 ids.ID, arg1 ids.NodeID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains.
func (mr *MockManagerMockRecorder) Contains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockManager)(nil).Contains), arg0, arg1)
}

// GetMap mocks base method.
func (m *MockManager) GetMap(arg0 ids.ID) map[ids.NodeID]*GetValidatorOutput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMap", arg0)
	ret0, _ := ret[0].(map[ids.NodeID]*GetValidatorOutput)
	return ret0
}

// GetMap indicates an expected call of GetMap.
func (mr *MockManagerMockRecorder) GetMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMap", reflect.TypeOf((*MockManager)(nil).GetMap), arg0)
}

// GetValidator mocks base method.
func (m *MockManager) GetValidator(arg0 ids.ID, arg1 ids.NodeID) (*Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", arg0, arg1)
	ret0, _ := ret[0].(*Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockManagerMockRecorder) GetValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockManager)(nil).GetValidator), arg0, arg1)
}

// GetValidatorIDs mocks base method.
func (m *MockManager) GetValidatorIDs(arg0 ids.ID) ([]ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorIDs", arg0)
	ret0, _ := ret[0].([]ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorIDs indicates an expected call of GetValidatorIDs.
func (mr *MockManagerMockRecorder) GetValidatorIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorIDs", reflect.TypeOf((*MockManager)(nil).GetValidatorIDs), arg0)
}

// GetWeight mocks base method.
func (m *MockManager) GetWeight(arg0 ids.ID, arg1 ids.NodeID) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeight", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWeight indicates an expected call of GetWeight.
func (mr *MockManagerMockRecorder) GetWeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeight", reflect.TypeOf((*MockManager)(nil).GetWeight), arg0, arg1)
}

// Len mocks base method.
func (m *MockManager) Len(arg0 ids.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockManagerMockRecorder) Len(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockManager)(nil).Len), arg0)
}

// RegisterCallbackListener mocks base method.
func (m *MockManager) RegisterCallbackListener(arg0 ids.ID, arg1 SetCallbackListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterCallbackListener", arg0, arg1)
}

// RegisterCallbackListener indicates an expected call of RegisterCallbackListener.
func (mr *MockManagerMockRecorder) RegisterCallbackListener(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCallbackListener", reflect.TypeOf((*MockManager)(nil).RegisterCallbackListener), arg0, arg1)
}

// RemoveWeight mocks base method.
func (m *MockManager) RemoveWeight(arg0 ids.ID, arg1 ids.NodeID, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWeight", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWeight indicates an expected call of RemoveWeight.
func (mr *MockManagerMockRecorder) RemoveWeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWeight", reflect.TypeOf((*MockManager)(nil).RemoveWeight), arg0, arg1, arg2)
}

// Sample mocks base method.
func (m *MockManager) Sample(arg0 ids.ID, arg1 int) ([]ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sample", arg0, arg1)
	ret0, _ := ret[0].([]ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sample indicates an expected call of Sample.
func (mr *MockManagerMockRecorder) Sample(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sample", reflect.TypeOf((*MockManager)(nil).Sample), arg0, arg1)
}

// String mocks base method.
func (m *MockManager) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockManagerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockManager)(nil).String))
}

// SubsetWeight mocks base method.
func (m *MockManager) SubsetWeight(arg0 ids.ID, arg1 set.Set[ids.NodeID]) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubsetWeight", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubsetWeight indicates an expected call of SubsetWeight.
func (mr *MockManagerMockRecorder) SubsetWeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubsetWeight", reflect.TypeOf((*MockManager)(nil).SubsetWeight), arg0, arg1)
}

// TotalWeight mocks base method.
func (m *MockManager) TotalWeight(arg0 ids.ID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalWeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalWeight indicates an expected call of TotalWeight.
func (mr *MockManagerMockRecorder) TotalWeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalWeight", reflect.TypeOf((*MockManager)(nil).TotalWeight), arg0)
}
