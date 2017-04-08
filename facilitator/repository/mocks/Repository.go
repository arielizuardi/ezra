package mocks

import "github.com/stretchr/testify/mock"

import "github.com/arielizuardi/ezra/facilitator"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetFacilitator provides a mock function with given fields: facilitatorID
func (_m *Repository) GetFacilitator(facilitatorID int64) (*facilitator.Facilitator, error) {
	ret := _m.Called(facilitatorID)

	var r0 *facilitator.Facilitator
	if rf, ok := ret.Get(0).(func(int64) *facilitator.Facilitator); ok {
		r0 = rf(facilitatorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*facilitator.Facilitator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(facilitatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
