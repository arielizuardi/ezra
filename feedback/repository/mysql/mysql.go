package mysql

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/feedback"
)

// MySQLFeedbackRepository ...
type MySQLFeedbackRepository struct {
	DBConn *sql.DB
}

// FetchAllFeedbackFields ...
func (m *MySQLFeedbackRepository) FetchAllFeedbackFields() ([]*feedback.Field, error) {

	rows, err := m.DBConn.Query(`SELECT id, name, description FROM feedback_field`)
	if err != nil {
		return nil, err
	}

	var fields []*feedback.Field
	for rows.Next() {
		f := new(feedback.Field)
		err := rows.Scan(&f.ID, &f.Name, &f.Description)
		if err != nil {
			return nil, err
		}

		fields = append(fields, f)
	}

	return fields, nil
}

// FetchFacilitatorFeedbacks ...
func (m *MySQLFeedbackRepository) FetchFacilitatorFeedbacks(facilitatorID int64, c *class.Class) ([]*feedback.FacilitatorFeedback, error) {
	return nil, nil
}

// FetchPresenterFeedbacks ...
func (m *MySQLFeedbackRepository) FetchPresenterFeedbacks(presenterID int64, c *class.Class, s *class.Session) ([]*feedback.PresenterFeedback, error) {
	return nil, nil
}

// StorePresenterFeedbacks ...
func (m *MySQLFeedbackRepository) StorePresenterFeedbacks(feedbacks []*feedback.PresenterFeedback) error {

	// `fields` should be in JSON format
	feedbackQuery := `INSERT INTO feedback_presenter (class_id, session_id, presenter_id, participant_email, fields, created_at, updated_at) ` +
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

		_, err = feedbackStmt.Exec(fd.Class.ID, fd.Session.ID, fd.Presenter.ID, fd.Participant.Email, string(byteFields), now, now)
		if err != nil {
			trx.Rollback()
			return err
		}
	}

	trx.Commit()

	return nil
}

func NewMySQLFeedbackRepository(dbConn *sql.DB) *MySQLFeedbackRepository {
	return &MySQLFeedbackRepository{DBConn: dbConn}
}
