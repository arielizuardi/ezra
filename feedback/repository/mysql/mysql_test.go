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
		t.Skip("Skip placement repository test")
	}

	suite.Run(t, new(MySQLTest))
}

func (s *MySQLTest) SetupTest() {
	errs, ok := db.RunAllMigrations(s.DSN)
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
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
	pt.ID = int64(1)

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
