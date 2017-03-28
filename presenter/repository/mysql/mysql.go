package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/arielizuardi/ezra/presenter"
)

type MySQLPresenterRepository struct {
	DBConn *sql.DB
}

func (m *MySQLPresenterRepository) GetPresenter(presenterID int64) (*presenter.Presenter, error) {
	row := m.DBConn.QueryRow(`SELECT id, name, description, profile_picture FROM presenter WHERE id = ? `, presenterID)

	p := new(presenter.Presenter)
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.ProfilePicture)
	if err != nil {
		fmt.Printf(`%v`, err)
		return nil, err
	}

	return p, nil
}

func (m *MySQLPresenterRepository) StorePresenter(p *presenter.Presenter) error {
	now := time.Now()
	res, err := m.DBConn.Exec(`INSERT INTO presenter (name, description, profile_picture, created_at, updated_at) VALUE (?, ?, ?, ?, ?)`, p.Name, p.Description, p.ProfilePicture, now, now)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = lastID

	return nil
}

func (m *MySQLPresenterRepository) FetchAllPresenters() ([]*presenter.Presenter, error) {
	res, err := m.DBConn.Query(`SELECT id, name, description, profile_picture FROM presenter`)
	if err != nil {
		return nil, err
	}

	var presenters []*presenter.Presenter
	if res.Next() {
		p := new(presenter.Presenter)
		err := res.Scan(&p.ID, &p.Name, &p.Description, &p.ProfilePicture)
		if err != nil {
			return nil, err
		}
		presenters = append(presenters, p)
	}

	return presenters, nil
}

func NewMySQLPresenterRepository(dbConn *sql.DB) *MySQLPresenterRepository {
	return &MySQLPresenterRepository{DBConn: dbConn}
}
