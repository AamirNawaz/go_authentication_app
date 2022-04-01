package db

import (
	"go_authentication_app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//creat db refrence to get it accessable on conroller
var DB *gorm.DB

func Connect() {

	dbConnection, err := gorm.Open(mysql.Open("aamir:aamir@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("connection failed to the database ")
	}

	DB = dbConnection
	//connection.Debug().AutoMigrate(&models.User{})
	//UpdateTableSchema(dbConnection)

}

//UpdateTableSchema for user table migration
func UpdateTableSchema(connection *gorm.DB) {
	connection.Debug().AutoMigrate(&models.User{})
}
