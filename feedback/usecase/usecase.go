package usecase

import (
	"strconv"
	"time"

	c "github.com/arielizuardi/ezra/class/repository"
	fcl "github.com/arielizuardi/ezra/facilitator/repository"
	"github.com/arielizuardi/ezra/feedback"
	f "github.com/arielizuardi/ezra/feedback/repository"
	"github.com/arielizuardi/ezra/participant"
	prt "github.com/arielizuardi/ezra/participant/repository"
	p "github.com/arielizuardi/ezra/presenter/repository"
)

var (
	ParticipantNameFieldID = 2
)

type Mapping struct {
	HeaderID int64 `json:"header_id"`
	FieldID  int64 `json:"field_id"`
}

type FeedbackUsecase interface {
	FetchAllFeedbackFields() ([]*feedback.Field, error)
	StorePresenterFeedbackWithMapping(presenterID int64, classID string, sessionID int64, mappings []*Mapping, values [][]string) ([]*feedback.PresenterFeedback, error)
}

type feedbackUsecase struct {
	ClassRepository       c.Repository
	PresenterRepository   p.Repository
	FacilitatorRepository fcl.Repository
	ParticipantRepository prt.Repository
	FeedbackRepository    f.Repository
}

func (u *feedbackUsecase) FetchAllFeedbackFields() ([]*feedback.Field, error) {
	return u.FetchAllFeedbackFields()
}

func (u *feedbackUsecase) StorePresenterFeedbackWithMapping(presenterID int64, classID string, sessionID int64, mappings []*Mapping, values [][]string) ([]*feedback.PresenterFeedback, error) {

	class, err := u.ClassRepository.GetClass(classID)
	if err != nil {
		return nil, err
	}

	presenter, err := u.PresenterRepository.GetPresenter(presenterID)
	if err != nil {
		return nil, err
	}

	var presenterFeedbacks []*feedback.PresenterFeedback
	// start loop from values, every loop represent every row
	for _, value := range values {
		pf := new(feedback.PresenterFeedback)
		pf.Class = class
		pf.Presenter = presenter
		ptc, fields, err := u.ConvertToParticipantAndFields(mappings, value)

		if err != nil {
			return nil, err
		}

		pf.Participant = ptc
		pf.Fields = fields

		presenterFeedbacks = append(presenterFeedbacks, pf)
	}

	if len(presenterFeedbacks) > 0 {
		if err := u.FeedbackRepository.StorePresenterFeedbacks(presenterFeedbacks); err != nil {
			return nil, err
		}
	}

	return presenterFeedbacks, nil
}

func (u *feedbackUsecase) ConvertToParticipantAndFields(mappings []*Mapping, value []string) (*participant.Participant, []*feedback.Field, error) {

	var fields []*feedback.Field
	var ptc *participant.Participant
	var err error

	// now we loop every column in value, find the match mapping then store to fields
	for headeridx, v := range value {
		for _, mapping := range mappings {
			mappingHeaderID := int(mapping.HeaderID)
			mappingFieldID := int(mapping.FieldID)

			if headeridx == mappingHeaderID && mappingFieldID == ParticipantNameFieldID {

				ptc, err = u.findOrCreateParticipant(v)
				if err != nil {
					return nil, nil, err
				}

			} else if headeridx == mappingHeaderID {
				fields = append(fields, &feedback.Field{ID: mapping.FieldID, Value: v})
			}
		}

	}

	return ptc, fields, nil
}

func (u *feedbackUsecase) findOrCreateParticipant(name string) (*participant.Participant, error) {
	p, err := u.ParticipantRepository.GetParticipantByName(name)
	if err != nil {
		return nil, err
	}

	if p == nil {
		now := time.Now()
		unixnano := now.UnixNano()
		email := strconv.Itoa(int(unixnano)) + `@noemail.com`
		newParticipant := new(participant.Participant)
		newParticipant.Email = email
		newParticipant.Name = name

		err := u.ParticipantRepository.StoreParticipant(p)
		if err != nil {
			return nil, err
		}

		return newParticipant, nil
	}

	return p, nil
}

func NewFeedbackUsecase(
	classRepository c.Repository,
	presenterRepository p.Repository,
	facilitatorRepository fcl.Repository,
	participantRepository prt.Repository,
	feedbackRepository f.Repository,
) *feedbackUsecase {
	return &feedbackUsecase{
		classRepository,
		presenterRepository,
		facilitatorRepository,
		participantRepository,
		feedbackRepository,
	}
}
