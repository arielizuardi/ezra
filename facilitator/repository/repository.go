package repository

import "github.com/arielizuardi/ezra/facilitator"

type Repository interface {
	GetFacilitator(facilitatorID int64) (*facilitator.Facilitator, error)
	GetFacilitatorByName(name string) (*facilitator.Facilitator, error)
	StoreFacilitator(f *facilitator.Facilitator) error
}
