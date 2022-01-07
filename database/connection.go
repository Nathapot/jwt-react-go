package database

import (
	"github.com/Nathapot/jwt-react-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@/jwt_react_go"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
