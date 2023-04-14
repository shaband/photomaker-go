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
