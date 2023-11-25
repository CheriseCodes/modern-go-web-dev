package services

import (
	"net/http"
	"runners-postgresql/models"
	"runners-postgresql/repositories"
	"strconv"
)

type RunnersService struct {
	runnersRepository    *repositories.runnersRepository
	NewResultsRepository *repositories.ResultsRepository
}

func NewRunnersService(
	runnersRepository *repositories.RunnersRepository,
	resultsRepository *repositories.ResultsRepository) *RunnersService {
	return &RunnersService{
		runnersRepository: runnersRepository,
		resultsRepository: resultsRepository
	}
}

func (fs RunnersService) CreateRunner(
	runner *models.Runner
) (*models.Runner, *models.ResponseError) {
	// TODO
}

func (rs RunnersService) UpdateRunner(
	runner *models.Runner
) *models.ResponseError {
	responseErr := validateRunnerId(runner.ID)
	if responseErr != nil {
		return responseErr
	}
	responseErr := validateRunner(runner)
	if responseErr != nil {
		return responseErr
	}
	return rs.runnersRepository.UpdateRunner(runner)
}

func (rs RunnersService) DeleteRunner(
	runnerId string
) *models.ResponseError {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return responseErr
	}
	return rs.runnersRepository.DeleteRunner(runnerId)
}

func (rs RunnersService) GetRunner(
	runnerId string
) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	runner, responseErr := rs.runnersRepository.GetRunner(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	results, responseErr := rs.resultsRepository.GetAllRunnersResults(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	runner.Results = results
	return runner, nil
}

func (rs RunnersService) GetRunnersBatch(country string, year string) ([]*models.Runner, *models.ResponseError) {
	if country != "" && year != "" {
		return nil, &models.ResponseError{
			Message: "Only one parameter can be passed",
			Status: http.StatusBadRequest,
		}
	}
	// get the top 10 runners from one country
	if country != "" {
		return rs.runnersRepository.GetRunnersByCountry(country)
	}
	// get the top 10 runners of a given year
	if year != "" {
		intYear, err := strconv.Atoi(year)
		if err != nil {
			return &models.ResponseError{
				Message: "Invalid year",
				Status: http.StatusBadRequest,
			}
		}
		currentYear := time.Now().Year()
		if intYear < 0 || intYear > currentYear {
			return &models.ResponseError{
				Message: "Invalid year",
				Status: http.StatusBadRequest,
			}
		}
		return rs.runnersRepository.GetRunnersByYear(intYear)
	}
	// get all runners
	return rs.runnersRepository.GetAllRunners()
}

func validateRunner(
	runner *models.Runner
) *models.ResponseError {
    // validate name
	if runner.FirstName == "" {
		return &models.ResponseError{
			Message: "Invalid first name",
			Status: http.StatusBadRequest
		}
	}  
	if runner.LastName == "" {
		return &models.ResponseError{
			Message: "Invalid last name",
			Status: http.StatusBadRequest
		}
	}
	// validate 16 < age < 125
	if runner.Age <= 16 || runner.Age > 125 {
		return &models.ResponseError{
			Message: "Invalid age",
			Status: http.StatusBadRequest
		}
	}
	if runner.Country == "" {
		return &models.ResponseError{
			Message: "Invalid country",
			Status: http.StatusBadRequest
		}
	}
	return nil
}

func validateRunnerId(runnerId string) *models.ResponseError {
	if runnerId == "" {
		return &models.ResponseError{
			Message: "Invalid runner ID",
			Status: http.StatusBadRequest
		}
	}
	return nil
}

func CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunner(runner)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.runnersRepository.CreateRunner(runner)
}

