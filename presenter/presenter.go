package presenter

import "errors"

var (
	ErrPresenterNotFound = errors.New(`Presenter not found`)
)

type Presenter struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
}
