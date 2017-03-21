package mocks

import "github.com/stretchr/testify/mock"

import "github.com/arielizuardi/ezra/class"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetClass provides a mock function with given fields: classID
func (_m *Repository) GetClass(classID int64) (*class.Class, error) {
	ret := _m.Called(classID)

	var r0 *class.Class
	if rf, ok := ret.Get(0).(func(int64) *class.Class); ok {
		r0 = rf(classID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*class.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
