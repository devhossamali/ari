// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/devhossamali/ari/v5"
	mock "github.com/stretchr/testify/mock"
)

// Modules is an autogenerated mock type for the Modules type
type Modules struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *Modules) Data(key *ari.Key) (*ari.ModuleData, error) {
	ret := _m.Called(key)

	var r0 *ari.ModuleData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ModuleData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ModuleData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *Modules) Get(key *ari.Key) *ari.ModuleHandle {
	ret := _m.Called(key)

	var r0 *ari.ModuleHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ModuleHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ModuleHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: filter
func (_m *Modules) List(filter *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(filter)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(*ari.Key) []*ari.Key); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Load provides a mock function with given fields: key
func (_m *Modules) Load(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reload provides a mock function with given fields: key
func (_m *Modules) Reload(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unload provides a mock function with given fields: key
func (_m *Modules) Unload(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
