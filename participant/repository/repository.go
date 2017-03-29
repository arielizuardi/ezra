package repository

import "github.com/arielizuardi/ezra/participant"

// Repository ...
type Repository interface {
	GetParticipantByName(name string) (*participant.Participant, error)
	GetParticipant(participantID int64) (*participant.Participant, error)
	StoreParticipant(p *participant.Participant) error
}
