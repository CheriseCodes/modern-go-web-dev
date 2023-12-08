package repositories

import (
	"database/sql"
	"runners-postgresql/models"
)

// use interfaces to abstract the implementation of the database so any database can be used
type ResultsRepositoryInterface interface {
	CreateResult(result *models.Result) (*models.Result, *models.ResponseError)
	DeleteResult(resultId string) (*models.Result, *models.ResponseError)
	GetAllRunnersResults(runnerId string) ([]*models.Result, *models.ResponseError)
	GetPersonalBestResults(runnerId string) (string, *models.ResponseError)
	GetSeasonBestResults(runnerId string, year int) (string, *models.ResponseError)
	SetTransaction(transaction *sql.Tx)
	GetHandler() *sql.DB
	GetTransaction() *sql.Tx
}

type RunnersRepositoryInterface interface {
	CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError)
	UpdateRunner(runner *models.Runner) *models.ResponseError
	UpdateRunnerResults(runner *models.Runner) *models.ResponseError
	DeleteRunner(runnerId string) *models.ResponseError
	GetRunner(runnerId string) (*models.Runner, *models.ResponseError)
	GetAllRunners() ([]*models.Runner, *models.ResponseError)
	GetRunnersByCountry(country string) ([]*models.Runner, *models.ResponseError)
	GetRunnersByYear(year int) ([]*models.Runner, *models.ResponseError)
	SetTransaction(transaction *sql.Tx)
	GetHandler() *sql.DB
	GetTransaction() *sql.Tx
}
