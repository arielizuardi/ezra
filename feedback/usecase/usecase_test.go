package usecase_test

import (
	"errors"
	"testing"

	"github.com/arielizuardi/ezra/class"
	c "github.com/arielizuardi/ezra/class/repository/mocks"
	f "github.com/arielizuardi/ezra/facilitator/repository/mocks"
	fb "github.com/arielizuardi/ezra/feedback/repository/mocks"
	"github.com/arielizuardi/ezra/participant"
	p "github.com/arielizuardi/ezra/participant/repository/mocks"
	"github.com/arielizuardi/ezra/presenter"
	prt "github.com/arielizuardi/ezra/presenter/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/arielizuardi/ezra/feedback/usecase"
)

func TestStorePresenterFeedbackWithMapping(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(&class.Class{ID: `COL-B2-2016`}, nil)
	mockClassRepo.On(`GetSession`, sessionID).Return(&class.Session{ID: sessionID}, nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(&presenter.Presenter{Name: `Juferson Mangempis`}, nil)

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(&participant.Participant{Name: `Doe`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(nil)

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(nil)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.NoError(t, err)
	assert.Len(t, presenterFeedbacks, 3)

	assert.Equal(t, `COL-B2-2016`, presenterFeedbacks[0].Class.ID)
	assert.Equal(t, `Juferson Mangempis`, presenterFeedbacks[0].Presenter.Name)

	assert.Equal(t, `John`, presenterFeedbacks[0].Participant.Name)
	assert.Equal(t, `Doe`, presenterFeedbacks[1].Participant.Name)
	assert.Equal(t, `Katnis`, presenterFeedbacks[2].Participant.Name)

	assert.Equal(t, `1490279819`, presenterFeedbacks[0].Fields[0].Value)
	assert.Equal(t, `1490279820`, presenterFeedbacks[1].Fields[0].Value)
	assert.Equal(t, `1490279825`, presenterFeedbacks[2].Fields[0].Value)

	assert.Equal(t, `4`, presenterFeedbacks[0].Fields[1].Value)
	assert.Equal(t, `3`, presenterFeedbacks[1].Fields[1].Value)
	assert.Equal(t, `3`, presenterFeedbacks[2].Fields[1].Value)

	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 1)
	mockFeedbackRepo.AssertCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}

func TestStorePresenterFeedbackWithMappingAndGetErrorFromClass(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(nil, errors.New(`Whoops!`))

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(&presenter.Presenter{Name: `Juferson Mangempis`}, nil)

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(&participant.Participant{Name: `Doe`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(nil)

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(nil)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.Error(t, err)
	assert.Len(t, presenterFeedbacks, 0)

	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 0)
	mockFeedbackRepo.AssertNotCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}

func TestStorePresenterFeedbackWithMappingAndGetErrorFromPresenter(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(&class.Class{}, nil)
	mockClassRepo.On(`GetSession`, sessionID).Return(&class.Session{ID: sessionID}, nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(nil, errors.New(`Whoops!`))

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(&participant.Participant{Name: `Doe`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(nil)

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(nil)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.Error(t, err)
	assert.Len(t, presenterFeedbacks, 0)

	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 0)
	mockFeedbackRepo.AssertNotCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}

func TestStorePresenterFeedbackWithMappingAndGetErrorFromParticipant(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(&class.Class{}, nil)
	mockClassRepo.On(`GetSession`, sessionID).Return(&class.Session{ID: sessionID}, nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(&presenter.Presenter{}, nil)

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(nil, errors.New(`Whoops!`))
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(nil)

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(nil)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.Error(t, err)
	assert.Len(t, presenterFeedbacks, 0)

	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertNotCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 0)
	mockFeedbackRepo.AssertNotCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}

func TestStorePresenterFeedbackWithMappingAndGetErrorWhenStoreParticipant(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(&class.Class{}, nil)
	mockClassRepo.On(`GetSession`, sessionID).Return(&class.Session{ID: sessionID}, nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(&presenter.Presenter{}, nil)

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(&participant.Participant{Name: `Katnis`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(errors.New(`Whoops!`))

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(nil)

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.Error(t, err)
	assert.Len(t, presenterFeedbacks, 0)

	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 1)
	mockFeedbackRepo.AssertNotCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}

func TestStorePresenterFeedbackWithMappingAndGetErrorWhenStoreFeedbacks(t *testing.T) {
	presenterID := int64(0)
	classID := `col-b2-2016`
	sessionID := int64(1)
	mappings := []*usecase.Mapping{&usecase.Mapping{HeaderID: 0, FieldID: 1}, &usecase.Mapping{1, 2}, &usecase.Mapping{2, 4}}

	var values [][]string

	value1 := []string{`1490279819`, `John`, `4`}
	values = append(values, value1)

	value2 := []string{`1490279820`, `Doe`, `3`}
	values = append(values, value2)

	value3 := []string{`1490279825`, `Katnis`, `3`}
	values = append(values, value3)

	mockClassRepo := new(c.Repository)
	mockClassRepo.On(`GetClass`, classID).Return(&class.Class{ID: `COL-B2-2016`}, nil)
	mockClassRepo.On(`GetSession`, sessionID).Return(&class.Session{ID: sessionID}, nil)

	mockPresenterRepo := new(prt.Repository)
	mockPresenterRepo.On(`GetPresenter`, presenterID).Return(&presenter.Presenter{Name: `Juferson Mangempis`}, nil)

	mockFacilitatorRepo := new(f.Repository)

	mockParticipantRepo := new(p.Repository)
	mockParticipantRepo.On(`GetParticipantByName`, `John`).Return(&participant.Participant{Name: `John`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Doe`).Return(&participant.Participant{Name: `Doe`}, nil)
	mockParticipantRepo.On(`GetParticipantByName`, `Katnis`).Return(nil, nil)
	mockParticipantRepo.On(`StoreParticipant`, mock.AnythingOfType(`*participant.Participant`)).Return(nil)

	mockFeedbackRepo := new(fb.Repository)
	mockFeedbackRepo.On(`StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`)).Return(errors.New(`Whoops!`))

	u := usecase.NewFeedbackUsecase(
		mockClassRepo,
		mockPresenterRepo,
		mockFacilitatorRepo,
		mockParticipantRepo,
		mockFeedbackRepo,
	)

	presenterFeedbacks, err := u.StorePresenterFeedbackWithMapping(presenterID, classID, sessionID, mappings, values)
	assert.Error(t, err)
	assert.Len(t, presenterFeedbacks, 0)

	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `John`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Doe`)
	mockParticipantRepo.AssertCalled(t, `GetParticipantByName`, `Katnis`)
	mockParticipantRepo.AssertNumberOfCalls(t, `StoreParticipant`, 1)
	mockFeedbackRepo.AssertCalled(t, `StorePresenterFeedbacks`, mock.AnythingOfType(`[]*feedback.PresenterFeedback`))

}
