package usecase

import (
	"github.com/arielizuardi/ezra/presenter"
	"github.com/arielizuardi/ezra/presenter/repository"
)

type PresenterUsecase interface {
	FetchAllPresenters() ([]*presenter.Presenter, error)
}

type presenterUsecase struct {
	PresenterRepository repository.Repository
}

func (u *presenterUsecase) FetchAllPresenters() ([]*presenter.Presenter, error) {
	return u.PresenterRepository.FetchAllPresenters()
}

func NewPresenterUsecase(presenterRepository repository.Repository) PresenterUsecase {
	return &presenterUsecase{PresenterRepository: presenterRepository}
}
