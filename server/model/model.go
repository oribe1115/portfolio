package model

import (
	"fmt"
	os "os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

var (
	db *gorm.DB
)

func EstablishConnection() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USERNAME")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("MYSQL_PASSWORD")
	if dbPass == "" {
		dbPass = "password"
	}

	dbHost := os.Getenv("MYSQL_HOSTNAME")
	if dbHost == "" {
		dbHost = "db"
	}

	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		dbName = "portfolio"
	}

	_db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return nil, xerrors.Errorf("Can't Connect to DATABASE: %w", err)
	}
	db = _db
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	if os.Getenv("GORM_DEBUG") != "" {
		db = db.Debug()
	}

	return db, nil
}

func Migration() error {
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return err
	}

	foreignKeys := [][5]string{
		{"sub_categories", "main_category_id", "main_categories(id)", "CASCADE", "CASCADE"},
		{"tagged_contents", "tag_id", "tags(id)", "CASCADE", "CASCADE"},
		{"tagged_contents", "content_id", "contents(id)", "CASCADE", "CASCADE"},
		{"main_images", "content_id", "contents(id)", "CASCADE", "CASCADE"},
		{"sub_images", "content_id", "contents(id)", "CASCADE", "CASCADE"},
		{"contents", "category_id", "sub_categories(id)", "CASCADE", "CASCADE"},
	}

	for _, c := range foreignKeys {
		fmt.Println(c)
		if err := db.Table(c[0]).AddForeignKey(c[1], c[2], c[3], c[4]).Error; err != nil {
			return err
		}
	}

	return nil
}

func IsErrRecordNotFound(err error) bool {
	return xerrors.Is(err, gorm.ErrRecordNotFound)
}

var allTables = []interface{}{
	&MainCategory{},
	&SubCategory{},
	&Content{},
	&MainImage{},
	&SubImage{},
	&Tag{},
	&TaggedContent{},
	&User{},
}
