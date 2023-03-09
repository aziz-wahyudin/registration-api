package mysql

import (
	"fmt"
	"log"

	"github.com/aziz-wahyudin/registration-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Participant{})
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

type Participant struct {
	gorm.Model
	Name  string
	Phone string `gorm:"unique"`
	Age   int
	Photo string
}
