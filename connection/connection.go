package connection

import (
	"fmt"

	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB(path string) (*gorm.DB, error) {
	// get env
	env, err := helper.GetEnv(path)
	if err != nil {
		return nil, err
	}

	DB_Host := env["DATABASE_HOST"]
	DB_Name := env["DATABASE_NAME"]
	DB_Port := env["DATABASE_PORT"]
	DB_Username := env["DATABASE_USERNAME"]
	DB_Password := env["DATABASE_PASSWORD"]

	// set connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_Username, DB_Password, DB_Host, DB_Port, DB_Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	db.AutoMigrate(&user.User{})
	return db, nil
}
