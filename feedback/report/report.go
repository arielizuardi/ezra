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
	Class      *class.Class         `json:"class"`
	Presenter  *presenter.Presenter `json:"presenter"`
	AvgFields  map[string]float64   `json:"avg_fields"`
	OverallAvg float64              `json:"overall_avg"`
	Comment    string               `json:"comment"`
}
