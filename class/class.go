package class

import "errors"

const (
	// COB const for Community of Believers
	COB = `COB`
	// COL const for Community of Leaders
	COL = `COL`
	// COC const for Community of Councellors
	COC = `COC`
)

var (
	ErrClassNotFound   = errors.New(`Class not found`)
	ErrSessionNotFound = errors.New(`Session not found`)
)

// Class represents class
type Class struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Batch int64  `json:"batch"`
	Year  int64  `json:"year"`
}

// Session represents session
type Session struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
