package report

import (
	"github.com/arielizuardi/ezra/facilitator"
	"github.com/arielizuardi/ezra/presenter"
)

type FacilitatorReport struct {
	Facilitator *facilitator.Facilitator
	AvgFields   map[string]float64
}

type PresenterReport struct {
	Presenter *presenter.Presenter
	AvgFields map[string]float64
}
