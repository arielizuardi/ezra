package usecase

import (
	"github.com/Sirupsen/logrus"
	c "github.com/arielizuardi/ezra/class/repository"
	fcl "github.com/arielizuardi/ezra/facilitator/repository"
	"github.com/arielizuardi/ezra/feedback"
	f "github.com/arielizuardi/ezra/feedback/repository"
	prt "github.com/arielizuardi/ezra/participant/repository"
	p "github.com/arielizuardi/ezra/presenter/repository"
)

type Mapping struct {
	HeaderID int64 `json:"header_id"`
	FieldID  int64 `json:"field_id"`
}

type FeedbackUsecase interface {
	FetchAllFeedbackFields() ([]*feedback.Field, error)
	StorePresenterFeedbackWithMapping(presenterID int64, classID string, sessionID int64, mappings []*Mapping, values [][]string) error
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

func (u *feedbackUsecase) StorePresenterFeedbackWithMapping(presenterID int64, classID string, sessionID int64, mappings []*Mapping, values [][]string) error {

	class, err := u.ClassRepository.GetClass(classID)
	if err != nil {
		return err
	}

	presenter, err := u.PresenterRepository.GetPresenter(presenterID)
	if err != nil {
		return err
	}

	var presenterFeedbacks []*feedback.PresenterFeedback
	// start loop from values, every loop represent every row
	for _, value := range values {

		pf := new(feedback.PresenterFeedback)
		pf.Class = class
		pf.Presenter = presenter

		participantID := int64(0) // must get from mapping information
		participant, err := u.ParticipantRepository.GetParticipant(participantID)
		if err != nil {
			logrus.Error(err)
			// just skip if no participant found for that particular id
			continue
		}

		pf.Participant = participant

		var fields []*feedback.Field
		// now we loop every column in value, find the match maapping then store to fields
		for headerIDX, v := range value {
			for _, mapping := range mappings {
				if headerIDX == int(mapping.HeaderID) {
					fields = append(fields, &feedback.Field{ID: mapping.FieldID, Value: v})
				}
			}
		}

		if len(fields) > 0 {
			pf.Fields = fields
		}

		presenterFeedbacks = append(presenterFeedbacks, pf)
	}

	if len(presenterFeedbacks) > 0 {
		if err := u.FeedbackRepository.StorePresenterFeedbacks(presenterFeedbacks); err != nil {
			return err
		}
	}

	// header {0:judul, 1:timestamp}
	// value [`judul`, 12345]
	// value [`judul2`, 12345]
	// rating {1: judul, 2: timestamp}
	// store(presenter_id, class_id, participant_id, rating_id, rating_value)
	// or store(presenter, class, participant, rating)
	// or store(presenter, class, participant, ratings)
	// if using no sql
	// {
	// 		"presenter": {},
	// 		"class" : {}
	//		"ratings": [{}, {}]
	// }
	// sql
	// presenter_id, class_id, rating_id, rating_value

	return nil
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
