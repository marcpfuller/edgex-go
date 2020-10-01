// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddDeviceProfile provides a mock function with given fields: e
func (_m *DBClient) AddDeviceProfile(e models.DeviceProfile) (models.DeviceProfile, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) models.DeviceProfile); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.DeviceProfile) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}