package mysql_test

import (
	"testing"

	"github.com/arielizuardi/ezra/db"
	"github.com/arielizuardi/ezra/participant"
	"github.com/arielizuardi/ezra/participant/repository/mysql"
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
	_, err := s.DBConn.Exec(`INSERT INTO participant (email, name, date, dob, phone_number) VALUE (?, ?, ?, ?, ?)`, `john@doe.com`, `John Doe`, `Epic`, `01-Jan-1990`, `+6281912001200`)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TestGetParticipantByName() {
	s.seed()
	r := mysql.NewMySQLParticipantRepository(s.DBConn)
	p, err := r.GetParticipantByName(`John Doe`)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), p)
	assert.Equal(s.T(), `John Doe`, p.Name)
}

func (s *MySQLTest) TestGetParticipant() {
	s.seed()
	r := mysql.NewMySQLParticipantRepository(s.DBConn)
	p, err := r.GetParticipant(`john@doe.com`)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), p)
	assert.Equal(s.T(), `John Doe`, p.Name)
}

func (s *MySQLTest) TestStoreParticipant() {
	r := mysql.NewMySQLParticipantRepository(s.DBConn)
	p := new(participant.Participant)
	p.Name = `Foo Bar`
	p.Email = `foo@bar.com`
	p.Date = `30-Jan-1990`
	p.PhoneNumber = `+6281912001200`
	err := r.StoreParticipant(p)
	assert.NoError(s.T(), err)
}
