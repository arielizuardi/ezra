package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Sirupsen/logrus"
	cfg "github.com/arielizuardi/ezra/config"

	classrepo "github.com/arielizuardi/ezra/class/repository/mysql"
	facilrepo "github.com/arielizuardi/ezra/facilitator/repository/mocks"
	feedbackrepo "github.com/arielizuardi/ezra/feedback/repository/mysql"
	participantrepo "github.com/arielizuardi/ezra/participant/repository/mysql"
	presenterrepo "github.com/arielizuardi/ezra/presenter/repository/mysql"

	feedbackusecase "github.com/arielizuardi/ezra/feedback/usecase"

	feedbackhttp "github.com/arielizuardi/ezra/feedback/http"

	reporthttp "github.com/arielizuardi/ezra/feedback/report/http"
	reportusecase "github.com/arielizuardi/ezra/feedback/report/usecase"

	classhttp "github.com/arielizuardi/ezra/class/http"
	classusecase "github.com/arielizuardi/ezra/class/usecase"

	presenterhttp "github.com/arielizuardi/ezra/presenter/http"
	presenterusecase "github.com/arielizuardi/ezra/presenter/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if config.GetBool(`debug`) {
		logrus.Warn(`Ezra is running in debug mode`)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

}

func main() {

	// Start the application
	// Setup MySQL Database
	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)

	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`

	logrus.Info(`Connecting to database`)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		logrus.Error(fmt.Sprintf(`Database connection failed. Err: %v`, err.Error()))
		os.Exit(1)
	}
	defer dbConn.Close()

	e := echo.New()

	classRepository := classrepo.NewMySQLClassRepository(dbConn)
	presenterRepository := presenterrepo.NewMySQLPresenterRepository(dbConn)
	facilitatorRepository := new(facilrepo.Repository)
	participantRepository := participantrepo.NewMySQLParticipantRepository(dbConn)
	feedbackRepository := feedbackrepo.NewMySQLFeedbackRepository(dbConn)

	classUsecase := classusecase.NewClassUsecase(classRepository)
	classhttp.Init(e, classUsecase)

	feedbackUsecase := feedbackusecase.NewFeedbackUsecase(classRepository, presenterRepository, facilitatorRepository, participantRepository, feedbackRepository)
	feedbackhttp.Init(e, feedbackUsecase)

	presenterUsecase := presenterusecase.NewPresenterUsecase(presenterRepository)
	presenterhttp.Init(e, presenterUsecase)

	// classRepository classRepository.Repository,
	// presenterRepository presenterRepository.Repository,
	// facilitatorRepository facilRepository.Repository,
	// feedbackRepository repository.Repository
	reportUsecase := reportusecase.NewReportUsecase(classRepository, presenterRepository, facilitatorRepository, feedbackRepository)
	reporthttp.Init(e, reportUsecase)

	e.Use(middleware.CORS())

	address := config.GetString(`server.address`)
	logrus.Infof(`Ezra server running at address : %v`, address)
	logrus.Fatal(e.Start(address))
}
