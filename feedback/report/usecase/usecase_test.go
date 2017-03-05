package usecase_test

import (
	"errors"
	"testing"

	"github.com/arielizuardi/ezra/facilitator"
	mocksFacilitatorRepository "github.com/arielizuardi/ezra/facilitator/repository/mocks"
	"github.com/arielizuardi/ezra/feedback"
	"github.com/arielizuardi/ezra/feedback/report/usecase"
	mocksFeedbackRepository "github.com/arielizuardi/ezra/feedback/repository/mocks"
	"github.com/arielizuardi/ezra/presenter"
	mocksPresenterRepository "github.com/arielizuardi/ezra/presenter/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePresenterReport(t *testing.T) {
	presenterID := int64(1)
	session := int64(1)
	batch := int64(1)
	year := int64(2017)

	p := new(presenter.Presenter)
	pr := new(mocksPresenterRepository.Repository)
	pr.On(`Get`, presenterID).Return(p, nil)

	r1 := feedback.NewRating(feedback.PresenterRatingKey[0], ``, 5)
	r2 := feedback.NewRating(feedback.PresenterRatingKey[1], ``, 4)
	r3 := feedback.NewRating(feedback.PresenterRatingKey[2], ``, 3)
	pf1 := new(feedback.PresenterFeedback)
	pf1.Ratings = []*feedback.Rating{r1, r2, r3}

	r4 := feedback.NewRating(feedback.PresenterRatingKey[0], ``, 3)
	r5 := feedback.NewRating(feedback.PresenterRatingKey[1], ``, 3)
	r6 := feedback.NewRating(feedback.PresenterRatingKey[2], ``, 3)
	pf2 := new(feedback.PresenterFeedback)
	pf2.Ratings = []*feedback.Rating{r4, r5, r6}

	r7 := feedback.NewRating(feedback.PresenterRatingKey[0], ``, 2)
	r8 := feedback.NewRating(feedback.PresenterRatingKey[1], ``, 1)
	r9 := feedback.NewRating(feedback.PresenterRatingKey[2], ``, 2)
	pf3 := new(feedback.PresenterFeedback)
	pf3.Ratings = []*feedback.Rating{r7, r8, r9}

	feedbacks := []*feedback.PresenterFeedback{pf1, pf2, pf3}
	fr := new(mocksFeedbackRepository.Repository)
	fr.On(`FetchPresenterFeedbacks`, presenterID, session, batch, year).Return(feedbacks, nil)

	fcr := new(mocksFacilitatorRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fr)

	report, err := u.GeneratePresenterReport(presenterID, session, batch, year)
	assert.NoError(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, report.AvgFields[feedback.PresenterRatingKey[0]], float64(10)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.PresenterRatingKey[1]], float64(8)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.PresenterRatingKey[2]], float64(8)/float64(3))

	pr.AssertCalled(t, `Get`, presenterID)
	fr.AssertCalled(t, `FetchPresenterFeedbacks`, presenterID, session, batch, year)
}

func TestGeneratePresenterReportNonExistsPresenter(t *testing.T) {
	presenterID := int64(1)
	session := int64(1)
	batch := int64(1)
	year := int64(2017)

	pr := new(mocksPresenterRepository.Repository)
	pr.On(`Get`, presenterID).Return(nil, nil)

	fr := new(mocksFeedbackRepository.Repository)
	fcr := new(mocksFacilitatorRepository.Repository)
	u := usecase.NewReportUsecase(pr, fcr, fr)

	report, err := u.GeneratePresenterReport(presenterID, session, batch, year)
	assert.EqualError(t, presenter.ErrPresenterNotFound, err.Error())
	assert.Nil(t, report)

	pr.AssertCalled(t, `Get`, presenterID)
	fr.AssertNotCalled(t, `FetchPresenterFeedbacks`, presenterID, session, batch, year)
}

func TestGeneratePresenterReportAndGetErrorFromPresenterRepository(t *testing.T) {
	presenterID := int64(1)
	session := int64(1)
	batch := int64(1)
	year := int64(2017)

	pr := new(mocksPresenterRepository.Repository)
	pr.On(`Get`, presenterID).Return(nil, errors.New(`Whoops!`))

	fr := new(mocksFeedbackRepository.Repository)
	fcr := new(mocksFacilitatorRepository.Repository)
	u := usecase.NewReportUsecase(pr, fcr, fr)

	report, err := u.GeneratePresenterReport(presenterID, session, batch, year)
	assert.EqualError(t, errors.New(`Whoops!`), err.Error())
	assert.Nil(t, report)

	pr.AssertCalled(t, `Get`, presenterID)
	fr.AssertNotCalled(t, `FetchPresenterFeedbacks`, presenterID, session, batch, year)
}

func TestGeneratePresenterReportAndGetErrorFromFeedbackRepository(t *testing.T) {
	presenterID := int64(1)
	session := int64(1)
	batch := int64(1)
	year := int64(2017)

	pr := new(mocksPresenterRepository.Repository)
	pr.On(`Get`, presenterID).Return(new(presenter.Presenter), nil)

	fr := new(mocksFeedbackRepository.Repository)
	fr.On(`FetchPresenterFeedbacks`, presenterID, session, batch, year).Return(nil, errors.New(`Whoops!`))

	fcr := new(mocksFacilitatorRepository.Repository)
	u := usecase.NewReportUsecase(pr, fcr, fr)

	report, err := u.GeneratePresenterReport(presenterID, session, batch, year)
	assert.EqualError(t, errors.New(`Whoops!`), err.Error())
	assert.Nil(t, report)

	pr.AssertCalled(t, `Get`, presenterID)
	fr.AssertCalled(t, `FetchPresenterFeedbacks`, presenterID, session, batch, year)
}

func TestGenerateFacilitatorReport(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	f := new(facilitator.Facilitator)
	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(f, nil)

	r1 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 5)
	r2 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 4)
	r3 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 3)
	pf1 := new(feedback.FacilitatorFeedback)
	pf1.Ratings = []*feedback.Rating{r1, r2, r3}

	r4 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 3)
	r5 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 3)
	r6 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 3)
	pf2 := new(feedback.FacilitatorFeedback)
	pf2.Ratings = []*feedback.Rating{r4, r5, r6}

	r7 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 2)
	r8 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 1)
	r9 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 2)
	pf3 := new(feedback.FacilitatorFeedback)
	pf3.Ratings = []*feedback.Rating{r7, r8, r9}

	feedbacks := []*feedback.FacilitatorFeedback{pf1, pf2, pf3}

	fbr := new(mocksFeedbackRepository.Repository)
	fbr.On(`FetchFacilitatorFeedbacks`, facilitatorID, batch, year).Return(feedbacks, nil)

	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.NoError(t, err)
	assert.NotNil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertCalled(t, `FetchFacilitatorFeedbacks`, facilitatorID, batch, year)

	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[0]], float64(10)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[1]], float64(8)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[2]], float64(8)/float64(3))
}

func TestGenerateFacilitatorReportNonExistsFacilitator(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(nil, nil)

	fbr := new(mocksFeedbackRepository.Repository)
	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.EqualError(t, facilitator.ErrFacilitatorNotFound, err.Error())
	assert.Nil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertNotCalled(t, `FetchFacilitatorFeedbacks`, facilitatorID, batch, year)
}

func TestGenerateFacilitatorReportAndGetErrorFromFacilitatorRepository(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(nil, errors.New(`Whoops!`))

	fbr := new(mocksFeedbackRepository.Repository)
	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.EqualError(t, errors.New(`Whoops!`), err.Error())
	assert.Nil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertNotCalled(t, `FetchFacilitatorFeedbacks`, facilitatorID, batch, year)
}

func TestGenerateFacilitatorReportAndGetErrorFromFeedbackRepository(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(new(facilitator.Facilitator), nil)

	fbr := new(mocksFeedbackRepository.Repository)
	fbr.On(`FetchFacilitatorFeedbacks`, facilitatorID, batch, year).Return(nil, errors.New(`Whoops!`))

	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.EqualError(t, errors.New(`Whoops!`), err.Error())
	assert.Nil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertCalled(t, `FetchFacilitatorFeedbacks`, facilitatorID, batch, year)
}
