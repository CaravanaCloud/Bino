package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	name := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to database \n%s", err.Error())
	}
}
