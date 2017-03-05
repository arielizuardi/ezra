package usecase

import (
	"github.com/arielizuardi/ezra/facilitator"
	"github.com/arielizuardi/ezra/feedback"
	"github.com/arielizuardi/ezra/feedback/report"
	"github.com/arielizuardi/ezra/feedback/repository"
	"github.com/arielizuardi/ezra/presenter"
	presenterRepository "github.com/arielizuardi/ezra/presenter/repository"

	facilRepository "github.com/arielizuardi/ezra/facilitator/repository"
)

// ReportUsecase defines report usecase
type ReportUsecase interface {
	GenerateFacilitatorReport(facilitatorID int64, batch int64, year int64) (*report.FacilitatorReport, error)
	GeneratePresenterReport(presenterID int64, session int64, batch int64, year int64) (*report.PresenterReport, error)
}

type reportUsecase struct {
	PresenterRepository   presenterRepository.Repository
	FacilitatorRepository facilRepository.Repository
	FeedbackRepository    repository.Repository
}

func (r *reportUsecase) GenerateFacilitatorReport(facilitatorID int64, batch int64, year int64) (*report.FacilitatorReport, error) {

	facil, err := r.FacilitatorRepository.Get(facilitatorID)
	if err != nil {
		return nil, err
	}

	if facil == nil {
		return nil, facilitator.ErrFacilitatorNotFound
	}

	facilitatorReport := new(report.FacilitatorReport)
	facilitatorReport.Facilitator = facil

	facilitatorFeedbacks, err := r.FeedbackRepository.FetchFacilitatorFeedbacks(facilitatorID, batch, year)
	if err != nil {
		return nil, err
	}

	sum := make(map[string]int64)
	ct := make(map[string]int64)
	avg := make(map[string]float64)

	for _, facilitatorFeedback := range facilitatorFeedbacks {
		for _, rating := range facilitatorFeedback.Ratings {
			_, ok := sum[rating.Key]
			if !ok {
				sum[rating.Key] = 0
			}

			_, ok2 := ct[rating.Key]
			if !ok2 {
				ct[rating.Key] = 0
			}

			sum[rating.Key] = sum[rating.Key] + rating.Score
			ct[rating.Key]++
		}
	}

	var sumAvg float64

	ctKey := 0
	for _, key := range feedback.FacilitatorRatingKey {
		_, ok3 := sum[key]
		if ok3 {
			avg[key] = float64(sum[key]) / float64(ct[key])
			sumAvg = sumAvg + avg[key]
			ctKey++
		}
	}

	facilitatorReport.AvgFields = avg

	if ctKey > 0 {
		facilitatorReport.OverallAvg = sumAvg / float64(ctKey)
	}

	return facilitatorReport, nil
}

func (r *reportUsecase) GeneratePresenterReport(presenterID int64, session int64, batch int64, year int64) (*report.PresenterReport, error) {

	p, err := r.PresenterRepository.Get(presenterID)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, presenter.ErrPresenterNotFound
	}

	presenterReport := new(report.PresenterReport)
	presenterReport.Presenter = p

	presenterFeedbacks, err := r.FeedbackRepository.FetchPresenterFeedbacks(presenterID, session, batch, year)
	if err != nil {
		return nil, err
	}

	sum := make(map[string]int64)
	ct := make(map[string]int64)
	avg := make(map[string]float64)

	for _, presenterFeedback := range presenterFeedbacks {
		for _, rating := range presenterFeedback.Ratings {
			_, ok := sum[rating.Key]
			if !ok {
				sum[rating.Key] = 0
			}

			_, ok2 := ct[rating.Key]
			if !ok2 {
				ct[rating.Key] = 0
			}

			sum[rating.Key] = sum[rating.Key] + rating.Score
			ct[rating.Key]++
		}
	}

	ctKey := 0
	var sumAvg float64
	for _, key := range feedback.PresenterRatingKey {
		_, ok3 := sum[key]
		if ok3 {
			avg[key] = float64(sum[key]) / float64(ct[key])
			sumAvg = sumAvg + avg[key]
			ctKey++
		}
	}

	presenterReport.AvgFields = avg

	if ctKey > 0 {
		presenterReport.OverallAvg = sumAvg / float64(ctKey)
	}

	return presenterReport, nil
}

// NewReportUsecase returns new instance of ReportUsecase
func NewReportUsecase(
	presenterRepository presenterRepository.Repository,
	facilitatorRepository facilRepository.Repository,
	feedbackRepository repository.Repository) ReportUsecase {

	return &reportUsecase{
		PresenterRepository:   presenterRepository,
		FacilitatorRepository: facilitatorRepository,
		FeedbackRepository:    feedbackRepository,
	}
}
