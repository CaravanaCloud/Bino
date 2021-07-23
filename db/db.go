package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var err error

  dsn := getDSN()
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database")
	}
}

// username:password@protocol(address)/dbname?param=value
// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func getDSN() string {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, name)
}
