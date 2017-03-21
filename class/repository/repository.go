package repository

import "github.com/arielizuardi/ezra/class"

type Repository interface {
	GetClass(classID int64) (*class.Class, error)
}
