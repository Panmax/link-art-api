package model

import (
	"fmt"
	"link-art-api/infrastructure/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func SaveOne(data interface{}) error {
	err := db.Save(data).Error
	return err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Account{})
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(config.DatabaseConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	migrate(db)
}
