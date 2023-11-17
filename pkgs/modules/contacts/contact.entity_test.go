package contacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestContactFake(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	contact := &Contact{}

	result := contact.Fake()

	assert.NotNil(t, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
