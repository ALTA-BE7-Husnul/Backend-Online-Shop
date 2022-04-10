package utils

import (
	"fmt"
	"group-project-2/configs"
	"group-project-2/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Cart{})
	db.AutoMigrate(&entities.Transaction{})
	db.AutoMigrate(&entities.Payment{})
	db.AutoMigrate(&entities.Address{})
	db.AutoMigrate(&entities.TransactionDetail{})
}
