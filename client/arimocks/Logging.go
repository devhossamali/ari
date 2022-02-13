// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/devhossamali/ari/v5"
	mock "github.com/stretchr/testify/mock"
)

// Logging is an autogenerated mock type for the Logging type
type Logging struct {
	mock.Mock
}

// Create provides a mock function with given fields: key, levels
func (_m *Logging) Create(key *ari.Key, levels string) (*ari.LogHandle, error) {
	ret := _m.Called(key, levels)

	var r0 *ari.LogHandle
	if rf, ok := ret.Get(0).(func(*ari.Key, string) *ari.LogHandle); ok {
		r0 = rf(key, levels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LogHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key, string) error); ok {
		r1 = rf(key, levels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Data provides a mock function with given fields: key
func (_m *Logging) Data(key *ari.Key) (*ari.LogData, error) {
	ret := _m.Called(key)

	var r0 *ari.LogData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.LogData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LogData)
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

// Delete provides a mock function with given fields: key
func (_m *Logging) Delete(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *Logging) Get(key *ari.Key) *ari.LogHandle {
	ret := _m.Called(key)

	var r0 *ari.LogHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.LogHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LogHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: filter
func (_m *Logging) List(filter *ari.Key) ([]*ari.Key, error) {
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

// Rotate provides a mock function with given fields: key
func (_m *Logging) Rotate(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
