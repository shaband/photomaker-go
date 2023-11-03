package categories

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *Service {

	return &Service{
		db: db,
	}
}

func (s Service) All(conds ...interface{}) []*Category {

	Categories := []*Category{}

	s.db.Find(&Categories, conds...)

	return Categories
}

func (s Service) GetSingleCategoryWithImages(conds ...interface{}) *Category {

	category := Category{}

	s.db.Preload("Images").Find(&category, conds)

	return &category
}

func (service *Service) GetAll() *[]Category {
	Categorys := []Category{}
	service.db.Find(&Categorys)
	return &Categorys
}

func (service *Service) Find(ID int) *Category {
	Category := Category{}
	service.db.Find(&Category, ID)

	return &Category
}
func (service Service) Update(ID uint, Category *Category) {
	Category.ID = ID
	_ = service.db.Save(Category)
}

func (service *Service) DeleteById(ID int) *Category {
	// Category:=service.Find(ID)
	service.db.Delete(&Category{}, ID)
	return nil
}

func (service *Service) Store(CategoryRequest *CreateCategoryRequest) *Category {

	Category := CategoryRequest.ToEntity()
	service.db.Create(Category)

	return Category
}
