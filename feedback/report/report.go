package report

import (
	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/facilitator"
	"github.com/arielizuardi/ezra/presenter"
)

type FacilitatorReport struct {
	Class       *class.Class
	Session     *class.Session
	Facilitator *facilitator.Facilitator
	AvgFields   map[string]float64
	OverallAvg  float64
	Summary     map[string]string
}

type PresenterReport struct {
	Class      *class.Class
	Presenter  *presenter.Presenter
	AvgFields  map[string]float64
	OverallAvg float64
	Comment    string
}
