package mysql_test

import (
	"testing"
	"time"

	"github.com/arielizuardi/ezra/db"
	"github.com/arielizuardi/ezra/presenter"
	"github.com/arielizuardi/ezra/presenter/repository/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MySQLTest struct {
	db.MySQLSuite
}

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip presenter repository test")
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
	_, err := s.DBConn.Exec(`INSERT INTO presenter (name, description, profile_picture, created_at, updated_at) VALUE (?, ?, ?, ?, ?)`, `Juferson Mangempis`, `This is description`, `http://lorempixel.com/100/100/people`, now, now)
	assert.NoError(s.T(), err)
}

func (s *MySQLTest) TestGetPresenter() {
	s.seed()
	r := mysql.NewMySQLPresenterRepository(s.DBConn)
	p, err := r.GetPresenter(int64(1))
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), p)
}

func (s *MySQLTest) TestFetchAllPresenters() {
	s.seed()
	r := mysql.NewMySQLPresenterRepository(s.DBConn)
	p, err := r.FetchAllPresenters()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), p, 1)
	assert.Equal(s.T(), `Juferson Mangempis`, p[0].Name)
}

func (s *MySQLTest) TestStorePresenter() {
	r := mysql.NewMySQLPresenterRepository(s.DBConn)
	p := new(presenter.Presenter)
	p.Name = `Yosie Martinus`
	p.Description = `Awesome description`
	err := r.StorePresenter(p)
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), p.ID)
}
