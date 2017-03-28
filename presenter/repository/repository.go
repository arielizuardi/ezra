package repository

import "github.com/arielizuardi/ezra/presenter"

type Repository interface {
	GetPresenter(presenterID int64) (*presenter.Presenter, error)
	FetchAllPresenters() ([]*presenter.Presenter, error)
	StorePresenter(p *presenter.Presenter) error
}
