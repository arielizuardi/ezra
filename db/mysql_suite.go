package db

import (
	"context"
	"database/sql"

	"github.com/arielizuardi/envase"
	"github.com/docker/docker/client"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MySQLSuite struct for MySQL Suite
type MySQLSuite struct {
	suite.Suite
	ContainerID string
	DSN         string
	DBConn      *sql.DB
}

// SetupSuite setup at the beginning of test
func (s *MySQLSuite) SetupSuite() {
	DisableLogging()

	var err error

	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	assert.NoError(s.T(), err)

	imageName := `mysql:5.7`
	host := `127.0.0.1`
	containerPort := `3306`
	exposedPort := `33060`
	containerName := `papua_test`

	user := `root`
	pass := `pass`
	databaseName := `jpcccol_db`
	envConfig := []string{
		`MYSQL_USER=` + user,
		`MYSQL_ROOT_PASSWORD=` + pass,
		`MYSQL_DATABASE=` + databaseName,
	}

	c := envase.NewDockerContainer(ctx, dockerClient, imageName, host, containerPort, exposedPort, containerName, envConfig)
	assert.NoError(s.T(), c.Start())
	s.DSN = user + `:` + pass + `@tcp(` + host + `:` + exposedPort + `)/` + databaseName

	s.DBConn, err = sql.Open("mysql", s.DSN+`?parseTime=true&loc=Asia%2FJakarta`)
	for {
		err := s.DBConn.Ping()
		if err == nil {
			break
		}
	}

	s.DBConn.Exec("set global sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")
	s.DBConn.Exec("set session sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';")
}

// TearDownSuite teardown at the end of test
func (s *MySQLSuite) TearDownSuite() {
	s.DBConn.Close()

	if s.ContainerID != `` {
		//StopContainer(s.ContainerID)
	}
}

func DisableLogging() {
	nopLogger := NopLogger{}
	mysql.SetLogger(nopLogger)
}

type NopLogger struct {
}

func (l NopLogger) Print(v ...interface{}) {
}
