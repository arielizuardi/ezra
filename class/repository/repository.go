package repository

import "github.com/arielizuardi/ezra/class"

type Repository interface {
	GetClass(classID string) (*class.Class, error)
}
