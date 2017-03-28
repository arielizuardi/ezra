package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/arielizuardi/ezra/class"
)

type MySQLClassRepository struct {
	DBConn *sql.DB
}

func (m *MySQLClassRepository) GetClass(classID string) (*class.Class, error) {
	row := m.DBConn.QueryRow(`SELECT id, name, batch, year FROM class WHERE id = ? `, classID)

	c := new(class.Class)
	err := row.Scan(&c.ID, &c.Name, &c.Batch, &c.Year)
	if err != nil {
		fmt.Printf(`%v`, err)
		return nil, err
	}

	return c, nil
}

func (m *MySQLClassRepository) StoreClass(c *class.Class) error {
	now := time.Now()
	batch := `B` + strconv.Itoa(int(c.Batch))
	year := strconv.Itoa(int(c.Year))
	c.ID = strings.Join([]string{c.Name, batch, year}, `-`)
	_, err := m.DBConn.Exec(`INSERT INTO class (id, name, batch, year, created_at, updated_at) VALUE (?, ?, ?, ?, ?, ?)`, c.ID, c.Name, c.Batch, c.Year, now, now)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLClassRepository) FetchAllClasses() ([]*class.Class, error) {
	res, err := m.DBConn.Query(`SELECT id, name, batch, year FROM class`)
	if err != nil {
		return nil, err
	}

	var classes []*class.Class
	if res.Next() {
		c := new(class.Class)
		err := res.Scan(&c.ID, &c.Name, &c.Batch, &c.Year)
		if err != nil {
			return nil, err
		}
		classes = append(classes, c)
	}

	return classes, nil
}

func NewMySQLClassRepository(dbConn *sql.DB) *MySQLClassRepository {
	return &MySQLClassRepository{DBConn: dbConn}
}
