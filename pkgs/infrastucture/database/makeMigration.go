package database

import (
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
	"gorm.io/gorm"
)

func MakeMigration(db *gorm.DB) {
	db.AutoMigrate(
		&categories.Category{},
		&categories.CategoryImage{},
		&clients.Client{},
		&contacts.Contact{},
		&services.Service{},
		&contacts.ServiceType{},
		&contacts.ServiceTypeItem{},
		&settings.Setting{},
		&sliders.Slider{},
		&users.User{},
	)
}
