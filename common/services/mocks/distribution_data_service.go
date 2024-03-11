// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	distribution "github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	mock "github.com/stretchr/testify/mock"
)

// DistributionDataService is an autogenerated mock type for the DistributionDataService type
type DistributionDataService struct {
	mock.Mock
}

// GetDistribution provides a mock function with given fields: root
func (_m *DistributionDataService) GetDistribution(root [32]byte) (*distribution.Distribution, error) {
	ret := _m.Called(root)

	if len(ret) == 0 {
		panic("no return value specified for GetDistribution")
	}

	var r0 *distribution.Distribution
	var r1 error
	if rf, ok := ret.Get(0).(func([32]byte) (*distribution.Distribution, error)); ok {
		return rf(root)
	}
	if rf, ok := ret.Get(0).(func([32]byte) *distribution.Distribution); ok {
		r0 = rf(root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*distribution.Distribution)
		}
	}

	if rf, ok := ret.Get(1).(func([32]byte) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDistribution provides a mock function with given fields: root, _a1
func (_m *DistributionDataService) SetDistribution(root [32]byte, _a1 *distribution.Distribution) error {
	ret := _m.Called(root, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SetDistribution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([32]byte, *distribution.Distribution) error); ok {
		r0 = rf(root, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDistributionDataService creates a new instance of DistributionDataService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDistributionDataService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DistributionDataService {
	mock := &DistributionDataService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}