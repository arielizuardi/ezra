package usecase_test

import (
	"testing"

	"github.com/arielizuardi/ezra/class"
	c "github.com/arielizuardi/ezra/class/repository/mocks"
	f "github.com/arielizuardi/ezra/facilitator/repository/mocks"
	fb "github.com/arielizuardi/ezra/feedback/repository/mocks"
	p "github.com/arielizuardi/ezra/participant/repository/mocks"
	"github.com/arielizuardi/ezra/presenter"
	prt "github.com/arielizuardi/ezra/presenter/repository/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/arielizuardi/ezra/feedback/usecase"
)

func TestStorePresenterFeedbackWithMapping(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{}
	values := [][]string{}

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(new(class.Class), nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(new(presenter.Presenter), nil)

	mockFacilitatorRepo := new(f.Repository)
	mockParticipantRepo := new(p.Repository)
	mockFeedbackRepo := new(fb.Repository)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.NoError(t, err)
}
