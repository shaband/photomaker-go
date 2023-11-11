package categories

import (
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}

type ServiceInterface interface {
	All(conds ...interface{}) []*Category
	GetSingleCategoryWithImages(conds ...interface{}) *Category
	GetAll() *[]Category
	Find(ID int) *Category
	Update(ctx *gin.Context, ID uint, CategoryRequest *UpdateCategoryRequest) *gorm.DB
	DeleteById(ID int) *gorm.DB
	Store(c *gin.Context, CategoryRequest *CreateCategoryRequest) *Category
	DeleteImageByCategoryId(id int) *gorm.DB
}
type Service struct {
	db DB
}

const filePath = "assets/uploads/categories"

func NewService(db *gorm.DB) *Service {

	return &Service{
		db: db,
	}
}

func (service Service) All(conds ...interface{}) []*Category {

	Categories := []*Category{}

	service.db.Find(&Categories, conds...)

	return Categories
}

func (service Service) GetSingleCategoryWithImages(conds ...interface{}) *Category {

	category := Category{}

	service.db.Preload("Images").Find(&category, conds)

	return &category
}

func (service *Service) GetAll() *[]Category {
	Categories := []Category{}
	service.db.Preload("Images").Find(&Categories)
	return &Categories
}

func (service *Service) Find(ID int) *Category {
	Category := Category{}
	service.db.Preload("Images").Find(&Category, ID)

	return &Category
}
func (service Service) Update(ctx *gin.Context, ID uint, CategoryRequest *UpdateCategoryRequest) *gorm.DB {
	category := CategoryRequest.ToEntity(ctx)
	category.ID = ID
	return service.db.Save(category)
}

func (service *Service) DeleteById(ID int) *gorm.DB {
	// Category:=service.Find(ID)
	return service.db.Delete(&Category{}, ID)

}

func (service *Service) Store(c *gin.Context, CategoryRequest *CreateCategoryRequest) *Category {
	category := CategoryRequest.ToEntity(c)
	service.db.Create(category)
	return category
}

func (service Service) DeleteImageByCategoryId(id int) *gorm.DB {
	categoryImage := &CategoryImage{}
	service.db.Find(categoryImage, id)
	return service.db.Delete(categoryImage)
}

func SaveImage(c *gin.Context, dest string, image *multipart.FileHeader) string {
	path := dest + "/" + fmt.Sprint(time.Now().Unix()) + "/" + image.Filename

	err := c.SaveUploadedFile(image, path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return "/" + path
}
func handleCategoryImages(ctx *gin.Context, images []*multipart.FileHeader) []CategoryImage {

	var categoryImages []CategoryImage
	for order, image := range images {
		categoryImages = append(
			categoryImages,
			CategoryImage{
				Order: uint(order),
				Image: SaveImage(ctx, filePath, image),
			})
	}
	return categoryImages
}
