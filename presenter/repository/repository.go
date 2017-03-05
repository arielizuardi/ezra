package repository

import "github.com/arielizuardi/ezra/presenter"

type Repository interface {
	Get(presenterID int64) (*presenter.Presenter, error)
}
