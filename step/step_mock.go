// Code generated by mockery v2.8.0. DO NOT EDIT.

package step

import (
	"fmt"
	"github.com/bitrise-io/go-utils/command"
	"github.com/stretchr/testify/mock"
	"os"
)

// mockCommandRunner is an autogenerated mock type for the commandRunner type
type mockCommandRunner struct {
	mock.Mock
}

// Run provides a mock function with given fields: model
func (_m *mockCommandRunner) Run(model *command.Model) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*command.Model) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunAndReturnExitCode provides a mock function with given fields: model
func (_m *mockCommandRunner) RunAndReturnExitCode(model *command.Model) (int, error) {
	ret := _m.Called(model)

	var r0 int
	if rf, ok := ret.Get(0).(func(*command.Model) int); ok {
		r0 = rf(model)
	} else {
		r0, ok = ret.Get(0).(int)
		if !ok {
			return 0, fmt.Errorf("cast error")
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*command.Model) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunAndReturnTrimmedOutput provides a mock function with given fields: model
func (_m *mockCommandRunner) RunAndReturnTrimmedOutput(model *command.Model) (string, error) {
	ret := _m.Called(model)

	var r0 string
	if rf, ok := ret.Get(0).(func(*command.Model) string); ok {
		r0 = rf(model)
	} else {
		r0, ok = ret.Get(0).(string)
		if !ok {
			return "", fmt.Errorf("cast error")
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*command.Model) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockTempDirProvider is an autogenerated mock type for the tempDirProvider type
type mockTempDirProvider struct {
	mock.Mock
}

// CreateTempDir provides a mock function with given fields: prefix
func (_m *mockTempDirProvider) CreateTempDir(prefix string) (string, error) {
	ret := _m.Called(prefix)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prefix)
	} else {
		r0, ok = ret.Get(0).(string)
		if !ok {
			return "", fmt.Errorf("cast error")
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockLogger is an autogenerated mock type for the logger type
type mockLogger struct {
	mock.Mock
}

// Debugf provides a mock function with given fields: format, v
func (_m *mockLogger) Debugf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Donef provides a mock function with given fields: format, v
func (_m *mockLogger) Donef(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, v
func (_m *mockLogger) Errorf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Printf provides a mock function with given fields: format, v
func (_m *mockLogger) Printf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Println provides a mock function with given fields:
func (_m *mockLogger) Println() {
	_m.Called()
}

// mockEnvManager is an autogenerated mock type for the envManager type
type mockEnvManager struct {
	mock.Mock
}

// Set provides a mock function with given fields: key, value
func (_m *mockEnvManager) Set(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsetByValue provides a mock function with given fields: value
func (_m *mockEnvManager) UnsetByValue(value string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFileWriter is an autogenerated mock type for the fileWriter type
type mockFileWriter struct {
	mock.Mock
}

// Write provides a mock function with given fields: path, value, mode
func (_m *mockFileWriter) Write(path string, value string, mode os.FileMode) error {
	ret := _m.Called(path, value, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, os.FileMode) error); ok {
		r0 = rf(path, value, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}