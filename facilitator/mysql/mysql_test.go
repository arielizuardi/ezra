package mysql_test

import (
	"testing"
	"time"

	"github.com/arielizuardi/ezra/db"
	"github.com/arielizuardi/ezra/facilitator"
	"github.com/arielizuardi/ezra/facilitator/repository/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MySQLTest struct {
	db.MySQLSuite
}

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip facilitator repository test")
	}

	suite.Run(t, new(MySQLTest))
}

func (s *MySQLTest) SetupTest() {
	errs, ok := db.RunAllMigrations(s.DSN)
	assert.True(s.T(), ok)
	assert.Len(s.T(), errs, 0)
}

func (s *MySQLTest) seed() {
	now := time.Now()
	_, err := s.DBConn.Exec(`INSERT INTO facilitator (name, description, profile_picture, created_at, updated_at) VALUE (?, ?, ?, ?, ?)`, `Arie Lizuardi`, `This is description`, `http://lorempixel.com/100/100/people`, now, now)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TestGetfacilitator() {
	s.seed()
	r := mysql.NewMySQLFacilitatorRepository(s.DBConn)
	p, err := r.GetFacilitator(int64(1))
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), p)
}

func (s *MySQLTest) TestStoreFacilitator() {
	r := mysql.NewMySQLFacilitatorRepository(s.DBConn)
	f := new(facilitator.Facilitator)
	f.Name = `Emma Watson`
	f.Description = `Awesome description`
	err := r.StoreFacilitator(f)
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), f.ID)
}
