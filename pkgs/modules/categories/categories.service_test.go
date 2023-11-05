package categories

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gdb, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gdb)

	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(sqlmock.NewRows([]string{"id", "name_ar", "name_en", "cover"}).AddRow(1, "test_ar", "test_en", "test_cover"))

	categories := service.All()

	assert.Equal(t, 1, len(categories))
	assert.Equal(t, "test_ar", categories[0].NameAr)
	assert.Equal(t, "test_en", categories[0].NameEn)
	assert.Equal(t, "test_cover", categories[0].Cover)
}

// Repeat similar tests for GetSingleCategoryWithImages, GetAll, Find, Update, DeleteById, Store, and DeleteImageByCategoryId
