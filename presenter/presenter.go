package presenter

import "errors"

var (
	ErrPresenterNotFound = errors.New(`Presenter not found`)
)

type Presenter struct {
}
