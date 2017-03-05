package repository

import "github.com/arielizuardi/ezra/facilitator"

type Repository interface {
	Get(facilitatorID int64) (*facilitator.Facilitator, error)
}
