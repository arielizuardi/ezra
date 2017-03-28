package mysql

import (
	"database/sql"
	"time"

	"github.com/arielizuardi/ezra/participant"
)

type MySQLParticipantRepository struct {
	DBConn *sql.DB
}

func (m *MySQLParticipantRepository) GetParticipant(email string) (*participant.Participant, error) {
	row := m.DBConn.QueryRow(`SELECT email, name, date, dob, phone_number FROM participant WHERE email = ? `, email)

	p := new(participant.Participant)
	err := row.Scan(&p.Email, &p.Name, &p.Date, &p.DOB, &p.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (m *MySQLParticipantRepository) GetParticipantByName(name string) (*participant.Participant, error) {
	row := m.DBConn.QueryRow(`SELECT email, name, date, dob, phone_number FROM participant WHERE name = ? `, name)

	p := new(participant.Participant)
	err := row.Scan(&p.Email, &p.Name, &p.Date, &p.DOB, &p.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (m *MySQLParticipantRepository) StoreParticipant(p *participant.Participant) error {
	now := time.Now()
	_, err := m.DBConn.Exec(`INSERT INTO participant (email, name, date, dob, phone_number, created_at, updated_at) VALUE (?, ?, ?, ?, ?, ?, ?)`, p.Email, p.Name, p.Date, p.DOB, p.PhoneNumber, now, now)
	if err != nil {
		return err
	}

	return nil
}

func NewMySQLParticipantRepository(dbConn *sql.DB) *MySQLParticipantRepository {
	return &MySQLParticipantRepository{DBConn: dbConn}
}
