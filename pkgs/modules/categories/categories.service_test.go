package categories

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

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

func TestService_GetAll(t *testing.T) {
	+	db, mock, _ := sqlmock.New()
	+	gormDB, _ := gorm.Open("postgres", db)
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)

	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
		AddRow(1, "NameAr1", "NameEn1", "Cover1").
		AddRow(2, "NameAr2", "NameEn2", "Cover2")

	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)

	categories := service.GetAll()

	assert.Len(t, categories, 2)
	+func TestService_Update(t *testing.T) {
	+	db, mock, _ := sqlmock.New()
	+	gormDB, _ := gorm.Open("postgres", db)
	+	service := NewCategoryService(gormDB)
	+	
	+	mock.ExpectExec("^UPDATE \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	+	
	+	categoryRequest := &UpdateCategoryRequest{NameAr: "UpdatedNameAr", NameEn: "UpdatedNameEn", Cover: "UpdatedCover"}
	+	service.Update(nil, 1, categoryRequest)
	+}
	+
	func TestService_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^UPDATE \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	categoryRequest := &UpdateCategoryRequest{NameAr: "UpdatedNameAr", NameEn: "UpdatedNameEn", Cover: "UpdatedCover"}
	service.Update(nil, 1, categoryRequest)
	}
	+func TestService_DeleteById(t *testing.T) {
	+	db, mock, _ := sqlmock.New()
	+	gormDB, _ := gorm.Open("postgres", db)
	+	service := NewCategoryService(gormDB)
	+	
	+	mock.ExpectExec("^DELETE FROM \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	+	
	+	service.DeleteById(1)
	+}
	+
	+func TestService_Store(t *testing.T) {
	+	db, mock, _ := sqlmock.New()
	+	gormDB, _ := gorm.Open("postgres", db)
	+	service := NewCategoryService(gormDB)
	+	
	+	mock.ExpectExec("^INSERT INTO \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	+	
	+	categoryRequest := &CreateCategoryRequest{NameAr: "NewNameAr", NameEn: "NewNameEn", Cover: "NewCover"}
	+	service.Store(nil, categoryRequest)
	+}
	+
	func TestService_Find(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
		AddRow(1, "NameAr1", "NameEn1", "Cover1")
	
	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)
	
	category := service.Find(1)
	
	assert.Equal(t, "NameAr1", category.NameAr)
	}
	
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^UPDATE \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	categoryRequest := &UpdateCategoryRequest{NameAr: "UpdatedNameAr", NameEn: "UpdatedNameEn", Cover: "UpdatedCover"}
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^DELETE FROM \"category_images\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.DeleteImageByCategoryId(1)
	}
	+
	+func TestService_Find(t *testing.T) {
	+	db, mock, _ := sqlmock.New()
	+	gormDB, _ := gorm.Open("postgres", db)
	+	service := NewCategoryService(gormDB)
	
	+	rows := sqlmock.NewRows([]string{"ID", "NameAr", "NameEn", "Cover"}).
	+		AddRow(1, "NameAr1", "NameEn1", "Cover1")
	+	
	+	mock.ExpectQuery("^SELECT (.+) FROM \"categories\"$").WillReturnRows(rows)
	+	
	+	category := service.Find(1)
	+	
	+	assert.Equal(t, "NameAr1", category.NameAr)
	}
	
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^UPDATE \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.Update(nil, 1, categoryRequest)
	
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^DELETE FROM \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.DeleteById(1)
	
	func TestService_Store(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^INSERT INTO \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	categoryRequest := &CreateCategoryRequest{NameAr: "NewNameAr", NameEn: "NewNameEn", Cover: "NewCover"}
	service.Store(nil, categoryRequest)
	}
	
	func TestService_DeleteImageByCategoryId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^DELETE FROM \"category_images\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.DeleteImageByCategoryId(1)
	}
	func TestService_DeleteById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^DELETE FROM \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.DeleteById(1)
	}
	
	func TestService_Store(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^INSERT INTO \"categories\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	categoryRequest := &CreateCategoryRequest{NameAr: "NewNameAr", NameEn: "NewNameEn", Cover: "NewCover"}
	service.Store(nil, categoryRequest)
	}
	
	func TestService_DeleteImageByCategoryId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	service := NewCategoryService(gormDB)
	
	mock.ExpectExec("^DELETE FROM \"category_images\"$").WillReturnResult(sqlmock.NewResult(1, 1))
	
	service.DeleteImageByCategoryId(1)
	}
