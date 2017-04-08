package repository

import "github.com/arielizuardi/ezra/participant"

// Repository ...
type Repository interface {
	GetParticipantByName(name string) (*participant.Participant, error)
	GetParticipant(email string) (*participant.Participant, error)
	StoreParticipant(p *participant.Participant) error
}
