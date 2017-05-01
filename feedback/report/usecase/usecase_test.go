package usecase_test

import (
	"errors"
	"testing"

	"github.com/arielizuardi/ezra/class"
	mockClass "github.com/arielizuardi/ezra/class/repository/mocks"
	mockFacil "github.com/arielizuardi/ezra/facilitator/repository/mocks"
	"github.com/arielizuardi/ezra/feedback"
	"github.com/arielizuardi/ezra/feedback/report/usecase"
	mockFeedback "github.com/arielizuardi/ezra/feedback/repository/mocks"
	"github.com/arielizuardi/ezra/presenter"
	mockPresenter "github.com/arielizuardi/ezra/presenter/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePresenterReport(t *testing.T) {
	presenterID := int64(1)
	classID := `jpcccol-b1-2016`
	sessionID := int64(1)

	classRepository := new(mockClass.Repository)
	c := &class.Class{ID: classID}
	classRepository.On(`GetClass`, classID).Return(c, nil)
	s := &class.Session{ID: sessionID}
	classRepository.On(`GetSession`, sessionID).Return(s, nil)

	presenterRepository := new(mockPresenter.Repository)
	p := &presenter.Presenter{ID: presenterID}
	presenterRepository.On(`GetPresenter`, presenterID).Return(p, nil)

	facilitatorRepository := new(mockFacil.Repository)

	feedbackRepository := new(mockFeedback.Repository)

	field1 := &feedback.Field{ID: int64(4), Name: `Penguasaan materi`, Value: `2`}
	field2 := &feedback.Field{ID: int64(5), Name: `Sistematika Penyajian`, Value: `3`}
	field3 := &feedback.Field{ID: int64(6), Name: `Gaya atau metode penyajian`, Value: `4`}
	field4 := &feedback.Field{ID: int64(7), Name: `Pengaturan Waktu`, Value: `3`}
	field5 := &feedback.Field{ID: int64(8), Name: `Penggunaan alat bantu`, Value: `4`}
	field6 := &feedback.Field{ID: int64(12), Name: `Another Filed`, Value: `Should not be count`}

	feedbackRepository.On("FetchAllFeedbackFields").Return([]*feedback.Field{field1, field2, field3, field4, field5, field6}, nil)

	p1 := &feedback.PresenterFeedback{Class: c, Session: s, Presenter: p, Fields: []*feedback.Field{field1, field2, field3, field4, field5, field6}}
	p2 := &feedback.PresenterFeedback{Class: c, Session: s, Presenter: p, Fields: []*feedback.Field{field1, field2, field3, field4, field5, field6}}
	p3 := &feedback.PresenterFeedback{Class: c, Session: s, Presenter: p, Fields: []*feedback.Field{field1, field2, field3, field4, field5, field6}}

	presenterFeedbacks := []*feedback.PresenterFeedback{p1, p2, p3}

	feedbackRepository.On(`FetchPresenterFeedbacks`, presenterID, c, s).Return(presenterFeedbacks, nil)

	u := usecase.NewReportUsecase(classRepository, presenterRepository, facilitatorRepository, feedbackRepository)

	report, err := u.GeneratePresenterReport(presenterID, classID, sessionID)
	assert.NoError(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, float64(2), report.AvgFields[`Penguasaan materi`])
	assert.Equal(t, float64(3), report.AvgFields[`Sistematika Penyajian`])
	assert.Equal(t, float64(4), report.AvgFields[`Gaya atau metode penyajian`])
	assert.Equal(t, float64(3), report.AvgFields[`Pengaturan Waktu`])
	assert.Equal(t, float64(4), report.AvgFields[`Penggunaan alat bantu`])

	assert.Equal(t, float64(3.2), report.OverallAvg)
}

func TestGeneratePresenterReportWithNotExistsClassAndNotExistsSession(t *testing.T) {
	presenterID := int64(1)
	classID := `jpcccol-b1-2016`
	sessionID := int64(1)

	classRepository := new(mockClass.Repository)
	classRepository.On(`GetClass`, classID).Return(nil, nil)

	presenterRepository := new(mockPresenter.Repository)
	facilitatorRepository := new(mockFacil.Repository)
	feedbackRepository := new(mockFeedback.Repository)
	u := usecase.NewReportUsecase(classRepository, presenterRepository, facilitatorRepository, feedbackRepository)
	_, err := u.GeneratePresenterReport(presenterID, classID, sessionID)
	assert.EqualError(t, class.ErrClassNotFound, err.Error())

	classRepository2 := new(mockClass.Repository)
	classRepository2.On(`GetClass`, classID).Return(&class.Class{}, nil)
	classRepository2.On(`GetSession`, sessionID).Return(nil, nil)

	u2 := usecase.NewReportUsecase(classRepository2, presenterRepository, facilitatorRepository, feedbackRepository)
	_, err = u2.GeneratePresenterReport(presenterID, classID, sessionID)
	assert.EqualError(t, class.ErrSessionNotFound, err.Error())
}

func TestGeneratePresenterReportNonExistsPresenter(t *testing.T) {
	presenterID := int64(1)
	classID := `jpcccol-b1-2016`
	sessionID := int64(1)

	classRepository := new(mockClass.Repository)
	classRepository.On(`GetClass`, classID).Return(&class.Class{}, nil)
	classRepository.On(`GetSession`, sessionID).Return(&class.Session{}, nil)

	presenterRepository := new(mockPresenter.Repository)
	presenterRepository.On(`GetPresenter`, presenterID).Return(nil, nil)

	facilitatorRepository := new(mockFacil.Repository)
	feedbackRepository := new(mockFeedback.Repository)
	u := usecase.NewReportUsecase(classRepository, presenterRepository, facilitatorRepository, feedbackRepository)
	_, err := u.GeneratePresenterReport(presenterID, classID, sessionID)
	assert.EqualError(t, presenter.ErrPresenterNotFound, err.Error())
}

func TestGeneratePresenterReportAndGetErrorFromFeedbackRepository(t *testing.T) {
	presenterID := int64(1)
	classID := `jpcccol-b1-2016`
	sessionID := int64(1)

	classRepository := new(mockClass.Repository)
	c := &class.Class{ID: classID}
	classRepository.On(`GetClass`, classID).Return(c, nil)
	s := &class.Session{ID: sessionID}
	classRepository.On(`GetSession`, sessionID).Return(s, nil)

	presenterRepository := new(mockPresenter.Repository)
	p := &presenter.Presenter{ID: presenterID}
	presenterRepository.On(`GetPresenter`, presenterID).Return(p, nil)

	facilitatorRepository := new(mockFacil.Repository)
	feedbackRepository := new(mockFeedback.Repository)
	feedbackRepository.On(`FetchPresenterFeedbacks`, presenterID, c, s).Return(nil, errors.New(`Whoops`))

	u := usecase.NewReportUsecase(classRepository, presenterRepository, facilitatorRepository, feedbackRepository)
	_, err := u.GeneratePresenterReport(presenterID, classID, sessionID)
	assert.Equal(t, `Whoops`, err.Error())
}

/*
func TestGenerateFacilitatorReport(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	f := new(facilitator.Facilitator)
	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(f, nil)

	fbc := new(feedback.Feedback)
	r1 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 5)
	r2 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 4)
	r3 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 3)
	fbc.Ratings = []*feedback.Rating{r1, r2, r3}

	fbc2 := new(feedback.Feedback)
	r4 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 3)
	r5 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 3)
	r6 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 3)

	fbc2.Ratings = []*feedback.Rating{r4, r5, r6}

	fbc3 := new(feedback.Feedback)
	r7 := feedback.NewRating(feedback.FacilitatorRatingKey[0], ``, 2)
	r8 := feedback.NewRating(feedback.FacilitatorRatingKey[1], ``, 1)
	r9 := feedback.NewRating(feedback.FacilitatorRatingKey[2], ``, 2)

	fbc3.Ratings = []*feedback.Rating{r7, r8, r9}

	feedbacks := []*feedback.Feedback{fbc, fbc2, fbc3}

	bff := new(feedback.BatchFeedbackFacilitator)
	bff.BagOfFeedback = feedbacks

	fbr := new(mocksFeedbackRepository.Repository)
	fbr.On(`GetBatchFeedbackFacilitator`, facilitatorID, batch, year).Return(bff, nil)

	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.NoError(t, err)
	assert.NotNil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertCalled(t, `GetBatchFeedbackFacilitator`, facilitatorID, batch, year)

	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[0]], float64(10)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[1]], float64(8)/float64(3))
	assert.Equal(t, report.AvgFields[feedback.FacilitatorRatingKey[2]], float64(8)/float64(3))
	assert.Equal(t, report.OverallAvg, 2.888888888888889)

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
	fbr.AssertNotCalled(t, `GetBatchFeedbackFacilitator`, facilitatorID, batch, year)
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
	fbr.AssertNotCalled(t, `GetBatchFeedbackFacilitator`, facilitatorID, batch, year)
}

func TestGenerateFacilitatorReportAndGetErrorFromFeedbackRepository(t *testing.T) {
	facilitatorID := int64(1)
	batch := int64(1)
	year := int64(2017)

	fcr := new(mocksFacilitatorRepository.Repository)
	fcr.On(`Get`, facilitatorID).Return(new(facilitator.Facilitator), nil)

	fbr := new(mocksFeedbackRepository.Repository)
	fbr.On(`GetBatchFeedbackFacilitator`, facilitatorID, batch, year).Return(nil, errors.New(`Whoops!`))

	pr := new(mocksPresenterRepository.Repository)

	u := usecase.NewReportUsecase(pr, fcr, fbr)

	report, err := u.GenerateFacilitatorReport(facilitatorID, batch, year)
	assert.EqualError(t, errors.New(`Whoops!`), err.Error())
	assert.Nil(t, report)

	fcr.AssertCalled(t, `Get`, facilitatorID)
	fbr.AssertCalled(t, `GetBatchFeedbackFacilitator`, facilitatorID, batch, year)
}
*/
