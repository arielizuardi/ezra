package usecase_test

import (
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

	feedbacks := []*feedback.PresenterFeedback{}
	fr := new(mocksFeedbackRepository.Repository)
	fr.On(`FetchPresenterFeedbacks`, presenterID, session, batch, year).Return(feedbacks, nil)

	fcr := new(mocksFacilitatorRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fr)

	report, err := u.GeneratePresenterReport(presenterID, session, batch, year)
	assert.NoError(t, err)
	assert.NotNil(t, report)

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

func TestGenerateFacilitatorReport(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	f := new(facilitator.Facilitator)
	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(f, nil)

	feedbacks := []*feedback.FacilitatorFeedback{}
	fbr := new(mocksFeedbackRepository.Repository)
	fbr.On(`FetchFacilitatorFeedbacks`, facilitatorID, batch, year).Return(feedbacks, nil)

	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.NoError(t, err)
	assert.NotNil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertCalled(t, `FetchFacilitatorFeedbacks`, facilitatorID, batch, year)
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
