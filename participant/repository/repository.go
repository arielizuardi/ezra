package repository

import "github.com/arielizuardi/ezra/participant"

type Repository interface {
	GetParticipant(participantID int64) (*participant.Participant, error)
}
