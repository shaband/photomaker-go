package categories

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestService_All(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)

	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
		AddRow(1, "NameAr1", "NameEn1", "Cover1").
		AddRow(2, "NameAr2", "NameEn2", "Cover2")

	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)

	categories := service.All()

	assert.Len(t, categories, 2)
	assert.Equal(t, "NameAr1", categories[0].NameAr)
	assert.Equal(t, "NameAr2", categories[1].NameAr)
}

// Similar test functions for other methods in the Service struct...
// Test for GetSingleCategoryWithImages method
func TestService_GetSingleCategoryWithImages(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)

	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
		AddRow(1, "NameAr1", "NameEn1", "Cover1")

	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)

	category := service.GetSingleCategoryWithImages()

	assert.Equal(t, "NameAr1", category.NameAr)
}

// Test for GetAll method
func TestService_GetAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)

	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
		AddRow(1, "NameAr1", "NameEn1", "Cover1").
		AddRow(2, "NameAr2", "NameEn2", "Cover2")

	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)

	categories := service.GetAll()

	assert.Len(t, categories, 2)
}

// Similar test functions for other methods in the Service struct...
