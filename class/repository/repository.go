package repository

import "github.com/arielizuardi/ezra/class"

type Repository interface {
	GetClass(classID string) (*class.Class, error)
	GetSession(sessionID int64) (*class.Session, error)
	FetchAllSessions() ([]*class.Session, error)
	FetchAllClasses() ([]*class.Class, error)
	StoreClass(c *class.Class) error
}
