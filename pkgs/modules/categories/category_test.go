package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestCategoryFake(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Create a Category instance
	category := &Category{}

	// Call the Fake function
	result := category.Fake()

	// Check that the Fake function returns a non-nil result
	assert.NotNil(t, result)

	// Check that the mock database's expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
