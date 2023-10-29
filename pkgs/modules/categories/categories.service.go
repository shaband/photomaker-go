package categories

import (
	"gorm.io/gorm"
)

type CategoriesService struct {
	db *gorm.DB
}

func NewContractService(db *gorm.DB) CategoriesService {

	return CategoriesService{
		db: db,
	}
}

func (s CategoriesService) All(conds ...interface{}) []*Category {

	Categories := []*Category{}

	s.db.Find(&Categories, conds...)

	return Categories
}

func (s CategoriesService) GetSingleCategoryWithImages(conds ...interface{}) *Category {

	category := Category{}

	s.db.Preload("Images").Find(&category, conds)

	return &category
}


func (service *CategoriesService) GetAll() *[]Category {
	Categorys := []Category{}
	service.db.Find(&Categorys)
	return &Categorys
}

func (service *CategoriesService) Find(ID int) *Category {
	Category := Category{}
	service.db.Find(&Category, ID)

	return &Category
}
func (service CategoriesService) Update(ID uint, Category *Category) {
	Category.ID = ID
	_ = service.db.Save(Category)
}

func (service *CategoriesService) DeleteById(ID int) *Category {
	// Category:=service.Find(ID)
	service.db.Delete(&Category{}, ID)
	return nil
}

func (service *CategoriesService) Store(CategoryRequest *CategoryRequest) *Category {

	Category := CategoryRequest.ToEntity()
	service.db.Create(Category)

	return Category
}