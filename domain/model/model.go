package model

import (
	"fmt"
	"link-art-api/infrastructure/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func SaveOne(data interface{}) error {
	err := DB.Save(data).Error
	return err
}

func CreateOne(data interface{}) error {
	err := DB.Create(data).Error
	return err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&FollowFlow{})
	db.AutoMigrate(&Approval{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Auction{})
	db.AutoMigrate(&Exhibition{})
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&MessageFlow{})
	db.AutoMigrate(&Province{})
	db.AutoMigrate(&City{})
	db.AutoMigrate(&County{})
}

// Setup initializes the database instance
func Setup() {
	var err error
	DB, err = gorm.Open(config.DatabaseConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(true)

	migrate(DB)
}
