package services

import (
	"net/http"
	"runners-postgresql/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRunnerInvalidFirstName(t *testing.T) {
	runner := &models.Runner{
		LastName: "Smith",
		Age:      30,
		Country:  "United States",
	}
	responseErr := validateRunner(runner)
	assert.NotEmpty(t, responseErr)
	assert.Equal(t, "Invalid first name", responseErr.Message)
	assert.Equal(t, http.StatusBadRequest, responseErr.Status)
}
