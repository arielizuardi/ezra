package facilitator

import "errors"

var (
	ErrFacilitatorNotFound = errors.New(`Facilitator not found`)
)

type Facilitator struct {
}
