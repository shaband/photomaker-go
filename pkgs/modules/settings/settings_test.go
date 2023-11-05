package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSettingFake(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	setting := &Setting{}

	result := setting.Fake()

	assert.NotNil(t, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
