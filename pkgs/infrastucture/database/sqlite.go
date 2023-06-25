package database

// import (
// 	_ "github.com/joho/godotenv/autoload"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"os"
// )

// type SqliteConnnection struct {
// 	dns string
// }

// func (connection SqliteConnnection) Connect(config *gorm.Config) (*gorm.DB, error) {
// 	return gorm.Open(sqlite.Open(connection.dns), config)
// }

// func NewsqliteConnection() SqliteConnnection {

// 	return SqliteConnnection{
// 		dns: "./" + os.Getenv("DB_NAME"),
// 	}
// }
