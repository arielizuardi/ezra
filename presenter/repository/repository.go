package repository

import "github.com/arielizuardi/ezra/presenter"

type Repository interface {
	GetPresenter(presenterID int64) (*presenter.Presenter, error)
}
