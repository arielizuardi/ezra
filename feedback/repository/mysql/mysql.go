package mysql

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/arielizuardi/ezra/feedback"
)

type MySQLFeedbackRepository struct {
	DBConn *sql.DB
}

func (m *MySQLFeedbackRepository) FetchFacilitatorFeedbacks(facilitatorID int64, batch int64, year int64) ([]*feedback.FacilitatorFeedback, error) {
	return nil, nil
}

func (m *MySQLFeedbackRepository) FetchPresenterFeedbacks(presenterID int64, session int64, batch int64, year int64) ([]*feedback.PresenterFeedback, error) {
	return nil, nil
}

func (m *MySQLFeedbackRepository) StorePresenterFeedbacks(feedbacks []*feedback.PresenterFeedback) error {

	// `fields` should be in JSON format
	feedbackQuery := `INSERT INTO feedback_presenter (class_id, session_id, presenter_id, participant_id, fields, created_at, updated_at) ` +
		` VALUES (?, ?, ?, ?, ?, ?, ?)`

	trx, err := m.DBConn.Begin()
	if err != nil {
		return err
	}

	feedbackStmt, err := trx.Prepare(feedbackQuery)
	if err != nil {
		return err
	}

	defer feedbackStmt.Close()

	for _, fd := range feedbacks {
		byteFields, err := json.Marshal(fd.Fields)
		if err != nil {
			trx.Rollback()
			return err
		}

		now := time.Now()

		_, err = feedbackStmt.Exec(fd.Class.ID, fd.Session.ID, fd.Presenter.ID, fd.Participant.ID, string(byteFields), now, now)
		if err != nil {
			trx.Rollback()
			return err
		}
	}

	trx.Commit()

	return nil
}
