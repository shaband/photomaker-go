package categories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"time"
)

type Service struct {
	db *gorm.DB
}

const filePath = "assets/uploads/categories"

func NewCategoryService(db *gorm.DB) *Service {

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
func (service Service) Update(ctx *gin.Context, ID uint, CategoryRequest *UpdateCategoryRequest) {
	category := CategoryRequest.ToEntity(ctx)
	category.ID = ID
	_ = service.db.Save(category)
}

func (service *Service) DeleteById(ID int) *Category {
	// Category:=service.Find(ID)
	service.db.Delete(&Category{}, ID)
	return nil
}

func (service *Service) Store(c *gin.Context, CategoryRequest *CreateCategoryRequest) *Category {
	category := CategoryRequest.ToEntity(c)
	service.db.Create(category)
	return category
}

func (service Service) DeleteImageByCategoryId(id int)  {
	categoryImage := &CategoryImage{}
	service.db.Find(categoryImage, id)
	// categoryID := categoryImage.CategoryId
	service.db.Delete(categoryImage)
	// return categoryID
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
