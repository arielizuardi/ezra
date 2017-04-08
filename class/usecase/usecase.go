package usecase

import (
	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/class/repository"
)

type ClassUsecase interface {
	FetchAllClasses() ([]*class.Class, error)
}

type classUsecase struct {
	ClassRepository repository.Repository
}

func (c *classUsecase) FetchAllClasses() ([]*class.Class, error) {
	return c.ClassRepository.FetchAllClasses()
}

func NewClassUsecase(classRepository repository.Repository) ClassUsecase {
	return &classUsecase{ClassRepository: classRepository}
}
