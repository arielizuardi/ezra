package mocks

import "github.com/stretchr/testify/mock"

import "github.com/arielizuardi/ezra/participant"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetParticipant provides a mock function with given fields: participantID
func (_m *Repository) GetParticipant(participantID int64) (*participant.Participant, error) {
	ret := _m.Called(participantID)

	var r0 *participant.Participant
	if rf, ok := ret.Get(0).(func(int64) *participant.Participant); ok {
		r0 = rf(participantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*participant.Participant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(participantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
