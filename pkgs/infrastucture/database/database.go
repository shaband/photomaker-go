package database

import (
	"fmt"

	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	DB, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
		&contacts.ServiceType{},
		&contacts.ServiceTypeItem{},
		&settings.Setting{},
		&sliders.Slider{},
	)
}
