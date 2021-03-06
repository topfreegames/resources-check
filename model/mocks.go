// Automatically generated by MockGen. DO NOT EDIT!
// Source: model/interface.go

package model

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of MonitorService interface
type MockMonitorService struct {
	ctrl     *gomock.Controller
	recorder *_MockMonitorServiceRecorder
}

// Recorder for MockMonitorService (not exported)
type _MockMonitorServiceRecorder struct {
	mock *MockMonitorService
}

func NewMockMonitorService(ctrl *gomock.Controller) *MockMonitorService {
	mock := &MockMonitorService{ctrl: ctrl}
	mock.recorder = &_MockMonitorServiceRecorder{mock}
	return mock
}

func (_m *MockMonitorService) EXPECT() *_MockMonitorServiceRecorder {
	return _m.recorder
}

func (_m *MockMonitorService) Send(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Send", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMonitorServiceRecorder) Send(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0...)
}

func (_m *MockMonitorService) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockMonitorServiceRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}
