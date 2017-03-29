package mysql_test

import (
	"testing"

	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/db"
	"github.com/arielizuardi/ezra/feedback"
	"github.com/arielizuardi/ezra/participant"
	"github.com/arielizuardi/ezra/presenter"

	"github.com/arielizuardi/ezra/feedback/repository/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MySQLTest struct {
	db.MySQLSuite
}

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip feedback repository test")
	}

	suite.Run(t, new(MySQLTest))
}

func (s *MySQLTest) SetupTest() {
	errs, ok := db.RunAllMigrations(s.DSN)
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
}

func (s *MySQLTest) seed() {
	_, err := s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Timestamp`, `Timestamp`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Nama Partisipan`, `Nama Partisipan`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `D.A.T.E`, `D.A.T.E`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Penguasaan Materi`, `Penguasaan Materi`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Sistematika Penyajian`, `Sistematika Penyajian`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Gaya atau metode penyajian`, `Gaya atau metode penyajian`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Pengaturan Waktu`, `Pengaturan Waktu`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Penggunaan alat bantu`, `Penggunaan alat bantu`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Nilai keseluruhan`, `Nilai keseluruhan`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Hal-hal yang saya suka`, `Hal-hal yang saya suka`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Hal-hal yang saya harapkan`, `Hal-hal yang saya harapkan`)
	assert.NoError(s.T(), err)
	_, err = s.DBConn.Exec(`INSERT INTO feedback_field (name, description) VALUE (?, ?)`, `Holy Discontent`, `Holy Discontent`)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TearDownTest() {
}

func (s *MySQLTest) TestStore() {
	m := &mysql.MySQLFeedbackRepository{s.DBConn}
	f := new(feedback.PresenterFeedback)
	c := new(class.Class)
	c.ID = `col-b2-2016`
	f.Class = c

	p := new(presenter.Presenter)
	p.ID = int64(1)
	f.Presenter = p

	ses := new(class.Session)
	ses.ID = int64(1)
	ses.Name = `Starting Point`
	f.Session = ses

	pt := new(participant.Participant)
	pt.Email = `john@doe.com`

	f.Participant = pt

	f1 := new(feedback.Field)
	f1.ID = int64(1)
	f1.Value = 3.0

	f2 := new(feedback.Field)
	f2.ID = int64(2)
	f2.Value = `Keren!`

	fields := []*feedback.Field{f1, f2}
	f.Fields = fields

	feedbacks := []*feedback.PresenterFeedback{f}
	err := m.StorePresenterFeedbacks(feedbacks)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TestFetchAllFeedbackFields() {
	s.seed()
	m := &mysql.MySQLFeedbackRepository{s.DBConn}
	res, err := m.FetchAllFeedbackFields()

	assert.NoError(s.T(), err)
	assert.Len(s.T(), res, 12)
}
