package usecase

import (
	"fmt"
	"strconv"

	"github.com/arielizuardi/ezra/class"
	classRepository "github.com/arielizuardi/ezra/class/repository"
	"github.com/arielizuardi/ezra/feedback/report"
	"github.com/arielizuardi/ezra/feedback/repository"
	"github.com/arielizuardi/ezra/presenter"
	presenterRepository "github.com/arielizuardi/ezra/presenter/repository"

	facilRepository "github.com/arielizuardi/ezra/facilitator/repository"
)

// ReportUsecase defines report usecase
type ReportUsecase interface {
	GenerateFacilitatorReport(facilitatorID int64, c *class.Class) (*report.FacilitatorReport, error)
	GeneratePresenterReport(presenterID int64, classID string, sessionID int64) (*report.PresenterReport, error)
}

type reportUsecase struct {
	ClassRepository       classRepository.Repository
	PresenterRepository   presenterRepository.Repository
	FacilitatorRepository facilRepository.Repository
	FeedbackRepository    repository.Repository
}

func (r *reportUsecase) GenerateFacilitatorReport(facilitatorID int64, c *class.Class) (*report.FacilitatorReport, error) {
	/*
		facil, err := r.FacilitatorRepository.Get(facilitatorID)
		if err != nil {
			return nil, err
		}

		if facil == nil {
			return nil, facilitator.ErrFacilitatorNotFound
		}

		facilitatorReport := new(report.FacilitatorReport)
		facilitatorReport.Facilitator = facil

		facilitatorFeedbacks, err := r.FeedbackRepository.FetchFacilitatorFeedbacks(facilitatorID, c)
		if err != nil {
			return nil, err
		}

		// sum := make(map[string]int64)
		// ct := make(map[string]int64)
		// avg := make(map[string]float64)

		for _, facilitatorFeedback := range facilitatorFeedbacks {
			for _, f := range facilitatorFeedback.Fields {
				switch f.ID {
				// TODO
				}
			}
		}

		// for _, facilitatorFeedback := range batchFeedbackFacilitator.BagOfFeedback {
		// 	for _, rating := range facilitatorFeedback.Ratings {
		// 		_, ok := sum[rating.Key]
		// 		if !ok {
		// 			sum[rating.Key] = 0
		// 		}
		//
		// 		_, ok2 := ct[rating.Key]
		// 		if !ok2 {
		// 			ct[rating.Key] = 0
		// 		}
		//
		// 		sum[rating.Key] = sum[rating.Key] + rating.Score
		// 		ct[rating.Key]++
		// 	}
		// }
		//
		// var sumAvg float64
		//
		// ctKey := 0
		// for _, key := range feedback.FacilitatorRatingKey {
		// 	_, ok3 := sum[key]
		// 	if ok3 {
		// 		avg[key] = float64(sum[key]) / float64(ct[key])
		// 		sumAvg = sumAvg + avg[key]
		// 		ctKey++
		// 	}
		// }

		// facilitatorReport.AvgFields = avg
		//
		// if ctKey > 0 {
		// 	facilitatorReport.OverallAvg = sumAvg / float64(ctKey)
		// }

		return facilitatorReport, nil*/

	return nil, nil
}

func (r *reportUsecase) GeneratePresenterReport(presenterID int64, classID string, sessionID int64) (*report.PresenterReport, error) {

	c, err := r.ClassRepository.GetClass(classID)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, class.ErrClassNotFound
	}

	s, err := r.ClassRepository.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	if s == nil {
		return nil, class.ErrSessionNotFound
	}

	p, err := r.PresenterRepository.GetPresenter(presenterID)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, presenter.ErrPresenterNotFound
	}

	presenterReport := new(report.PresenterReport)
	presenterReport.Presenter = p
	presenterReport.Class = c

	presenterFeedbacks, err := r.FeedbackRepository.FetchPresenterFeedbacks(presenterID, c, s)
	if err != nil {
		return nil, err
	}

	sum := make(map[string]int64)
	ct := make(map[string]int64)

	feedbackFields, err := r.FeedbackRepository.FetchAllFeedbackFields()
	if err != nil {
		return nil, err
	}

	fields := make(map[int64]string)
	for _, feedbackField := range feedbackFields {
		fields[feedbackField.ID] = feedbackField.Name
	}

	for _, presenterFeedback := range presenterFeedbacks {
		for _, f := range presenterFeedback.Fields {
			switch f.ID {
			case 4, 5, 6, 7, 8: // Penguasaan materi, Sistematika Penyajian, Gaya atau metode penyajian, Gaya atau metode penyajian, Pengaturan Waktu, Penggunaan alat bantu
				if f.Value != nil {
					val, err := strconv.Atoi(f.Value.(string))
					if err != nil {
						continue
					}

					name := fields[f.ID]
					sum[name] += int64(val)
					ct[name]++
				}
			default:
			}
		}
	}

	var keys []string
	for k := range sum {
		keys = append(keys, k)
	}

	var sumAvg float64
	ctKey := 0
	avg := make(map[string]float64)

	for _, key := range keys {
		_, ok := sum[key]
		if ok {
			avg[key] = float64(sum[key]) / float64(ct[key])
			sumAvg = sumAvg + avg[key]
			ctKey++
		}
	}

	fmt.Printf(`>>>> %v`, avg)

	presenterReport.AvgFields = avg

	if ctKey > 0 {
		presenterReport.OverallAvg = sumAvg / float64(ctKey)
	}

	return presenterReport, nil
}

// NewReportUsecase returns new instance of ReportUsecase
func NewReportUsecase(
	classRepository classRepository.Repository,
	presenterRepository presenterRepository.Repository,
	facilitatorRepository facilRepository.Repository,
	feedbackRepository repository.Repository) ReportUsecase {

	return &reportUsecase{
		ClassRepository:       classRepository,
		PresenterRepository:   presenterRepository,
		FacilitatorRepository: facilitatorRepository,
		FeedbackRepository:    feedbackRepository,
	}
}
