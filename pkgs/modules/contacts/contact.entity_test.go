package contacts

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestContactFake(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// contact := &Contact{}

	// result := contact.Fake()

	// assert.NotNil(t, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
