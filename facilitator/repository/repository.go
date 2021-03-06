package repository

import "github.com/arielizuardi/ezra/facilitator"

type Repository interface {
	GetFacilitator(facilitatorID int64) (*facilitator.Facilitator, error)
}
