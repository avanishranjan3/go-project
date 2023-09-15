package Controllers

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GormDb() (*gorm.DB, error) {
	// Database connection parameters
	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	// Open a database connection
	dsn := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")" + "/" + dbName
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
