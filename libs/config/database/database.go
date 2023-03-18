package database

import (
	"fmt"

	"github.com/shaband/photomaker-go/libs/modules/categories"
	"github.com/shaband/photomaker-go/libs/modules/clients"
	"github.com/shaband/photomaker-go/libs/modules/contacts"
	"github.com/shaband/photomaker-go/libs/modules/services"
	"github.com/shaband/photomaker-go/libs/modules/settings"
	"github.com/shaband/photomaker-go/libs/modules/sliders"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {

	DB, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(err)
	db = DB
}

func GetConnection() *gorm.DB {

	return db
}

func MakeMigration(db *gorm.DB) {

	db.AutoMigrate(
		&categories.Category{},
		&categories.CategoryImage{},
		&clients.Client{},
		&contacts.Contact{},
		&services.Service{},
		&contacts.Contact{},
		&contacts.ServiceType{},
		&settings.Setting{},
		&sliders.Slider{},
	)
}
