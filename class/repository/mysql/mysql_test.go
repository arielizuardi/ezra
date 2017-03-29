package mysql_test

import (
	"testing"
	"time"

	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/class/repository/mysql"
	"github.com/arielizuardi/ezra/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MySQLTest struct {
	db.MySQLSuite
}

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip class repository test")
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
	_, err := s.DBConn.Exec(`INSERT INTO class (id, name, batch, year, created_at, updated_at) VALUE (?, ?, ?, ?, ?, ?)`, `COL-B1-2016`, `COL`, int64(1), int64(2016), now, now)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TestStoreClass() {
	r := mysql.NewMySQLClassRepository(s.DBConn)
	c := new(class.Class)
	c.Name = class.COL
	c.Batch = 1
	c.Year = 2016

	assert.NoError(s.T(), r.StoreClass(c))
	assert.Equal(s.T(), `COL-B1-2016`, c.ID)
}

func (s *MySQLTest) TestGetClass() {
	s.seed()
	r := mysql.NewMySQLClassRepository(s.DBConn)
	c, err := r.GetClass(`COL-B1-2016`)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), c)
}

func (s *MySQLTest) TestFetchAllClasses() {
	s.seed()
	r := mysql.NewMySQLClassRepository(s.DBConn)
	classes, err := r.FetchAllClasses()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), classes, 1)
	assert.Equal(s.T(), `COL-B1-2016`, classes[0].ID)
}
