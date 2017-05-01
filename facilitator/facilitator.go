package facilitator

import "errors"

var (
	ErrFacilitatorNotFound = errors.New(`Facilitator not found`)
)

type Facilitator struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
}
