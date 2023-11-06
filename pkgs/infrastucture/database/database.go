package database

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type ConfigLoader interface {
	GetEnv(key string) string
}

func Init(configLoader ConfigLoader) {

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	connection, err := connectionFactory(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	DB, err := connection.Connect(config)
	if err != nil {
		panic("failed to connect database")
	}
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
		&users.User{},
	)
}



		type ConnectionFactory interface {
			Connect(config *gorm.Config) (*gorm.DB, error)
		}
		
		func connectionFactory(factory string) (ConnectionFactory, error) {
		
			if factory == "postgres" {
				return NewPostgresConnection(), nil
				// } else if factory == "sqlite" {
				// 	return NewsqliteConnection(), nil
			} else {
				return nil, errors.New("Invalid Connection")
			}
		}
	

