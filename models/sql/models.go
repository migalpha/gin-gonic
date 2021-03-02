package sql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                                     error
		dbType, dbName, user, password, host, port, tablePrefix string
	)

	dbType = os.Getenv("DB_TYPE")
	dbName = os.Getenv("DB_NAME")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	tablePrefix = ""

	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host,
		port,
		user,
		dbName,
		password))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		db.LogMode(true)
	}
	db.SingularTable(false)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
}

// CloseDB close connection
func CloseDB() {
	defer db.Close()
}
