package mysql

import (
	"database/sql"
	"time"

	"github.com/arielizuardi/ezra/facilitator"
)

type MySQLFacilitatorRepository struct {
	DBConn *sql.DB
}

func (r *MySQLFacilitatorRepository) GetFacilitator(facilitatorID int64) (*facilitator.Facilitator, error) {
	row := r.DBConn.QueryRow(`SELECT id, name, description, profile_picture FROM facilitator WHERE id = ? `, facilitatorID)
	f := new(facilitator.Facilitator)
	err := row.Scan(&f.ID, &f.Name, &f.Description, &f.ProfilePicture)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return f, nil
}

func (r *MySQLFacilitatorRepository) GetFacilitatorByName(name string) (*facilitator.Facilitator, error) {
	row := r.DBConn.QueryRow(`SELECT id, name, description, profile_picture FROM facilitator WHERE name = ? `, name)

	f := new(facilitator.Facilitator)
	err := row.Scan(&f.ID, &f.Name, &f.Description, &f.ProfilePicture)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return f, nil
}

func (r *MySQLFacilitatorRepository) StoreFacilitator(f *facilitator.Facilitator) error {
	now := time.Now()
	res, err := r.DBConn.Exec(
		`INSERT INTO facilitator (name, description, profile_picture, created_at, updated_at) VALUE (?, ?, ?, ?, ?)`,
		f.Name,
		f.Description,
		f.ProfilePicture,
		now,
		now,
	)

	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	f.ID = lastID

	return nil
}

func NewMySQLFacilitatorRepository(dbConn *sql.DB) *MySQLFacilitatorRepository {
	return &MySQLFacilitatorRepository{DBConn: dbConn}
}
