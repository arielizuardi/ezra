package usecase

import (
	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/class/repository"
)

type ClassUsecase interface {
	FetchAllClasses() ([]*class.Class, error)
	FetchAllSessions() ([]*class.Session, error)
}

type classUsecase struct {
	ClassRepository repository.Repository
}

func (c *classUsecase) FetchAllClasses() ([]*class.Class, error) {
	return c.ClassRepository.FetchAllClasses()
}

func (c *classUsecase) FetchAllSessions() ([]*class.Session, error) {
	return c.ClassRepository.FetchAllSessions()
}

func NewClassUsecase(classRepository repository.Repository) ClassUsecase {
	return &classUsecase{ClassRepository: classRepository}
}
