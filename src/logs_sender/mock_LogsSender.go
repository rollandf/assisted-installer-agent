// Code generated by mockery v1.0.0. DO NOT EDIT.

package logs_sender

import (
	strfmt "github.com/go-openapi/strfmt"
	mock "github.com/stretchr/testify/mock"
)

// MockLogsSender is an autogenerated mock type for the LogsSender type
type MockLogsSender struct {
	mock.Mock
}

// CreateFolderIfNotExist provides a mock function with given fields: folder
func (_m *MockLogsSender) CreateFolderIfNotExist(folder string) error {
	ret := _m.Called(folder)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Execute provides a mock function with given fields: command, args
func (_m *MockLogsSender) Execute(command string, args ...string) (string, string, int) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, ...string) string); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string, ...string) int); ok {
		r2 = rf(command, args...)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// ExecuteOutputToFile provides a mock function with given fields: outputFilePath, command, args
func (_m *MockLogsSender) ExecuteOutputToFile(outputFilePath string, command string, args ...string) (string, int) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, outputFilePath, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, ...string) string); ok {
		r0 = rf(outputFilePath, command, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, string, ...string) int); ok {
		r1 = rf(outputFilePath, command, args...)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// FileUploader provides a mock function with given fields: filePath, clusterID, hostID, inventoryUrl, pullSecretToken
func (_m *MockLogsSender) FileUploader(filePath string, clusterID strfmt.UUID, hostID strfmt.UUID, inventoryUrl string, pullSecretToken string) error {
	ret := _m.Called(filePath, clusterID, hostID, inventoryUrl, pullSecretToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, strfmt.UUID, strfmt.UUID, string, string) error); ok {
		r0 = rf(filePath, clusterID, hostID, inventoryUrl, pullSecretToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
