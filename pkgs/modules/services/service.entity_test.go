package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestServiceFake(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	service := &Service{}

	result := service.Fake()

	assert.NotNil(t, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
