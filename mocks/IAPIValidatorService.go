// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IAPIValidatorService is an autogenerated mock type for the IAPIValidatorService type
type IAPIValidatorService struct {
	mock.Mock
}

// ValidateStruct provides a mock function with given fields: req, name
func (_m *IAPIValidatorService) ValidateStruct(req interface{}, name string) (string, bool) {
	ret := _m.Called(req, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}, string) string); ok {
		r0 = rf(req, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(interface{}, string) bool); ok {
		r1 = rf(req, name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
