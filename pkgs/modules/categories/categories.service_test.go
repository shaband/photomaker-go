package categories

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestService_All(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	categories := service.All()

	assert.NotNil(t, categories)
}

func TestService_GetSingleCategoryWithImages(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	category := service.GetSingleCategoryWithImages()

	assert.NotNil(t, category)
}

func TestService_GetAll(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	categories := service.GetAll()

	assert.NotNil(t, categories)
}

func TestService_Find(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	category := service.Find(1)

	assert.NotNil(t, category)
}

func TestService_Update(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	categoryRequest := &UpdateCategoryRequest{}
	categoryRequest.Fake()

	service.Update(nil, 1, categoryRequest)

	category := service.Find(1)

	assert.NotNil(t, category)
}

func TestService_DeleteById(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	service.DeleteById(1)

	category := service.Find(1)

	assert.Nil(t, category)
}

func TestService_Store(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	categoryRequest := &CreateCategoryRequest{}
	categoryRequest.Fake()

	category := service.Store(nil, categoryRequest)

	assert.NotNil(t, category)
}

func TestService_DeleteImageByCategoryId(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	service := NewCategoryService(db)

	categoryID := service.DeleteImageByCategoryId(1)

	assert.Equal(t, 1, categoryID)
}




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

