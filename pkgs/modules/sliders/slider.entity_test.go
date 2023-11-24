package sliders

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSliderFake(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	slider := &Slider{}

	result := slider.Fake()

	assert.NotNil(t, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}
