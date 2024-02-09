package env

import "gorm.io/gorm"

var (
	Host     = "localhost"
	User     = "postgres"
	Password = "ravenclaw500"
	DbPort   = "5432"
	DbName   = "golang-bootcamp-2"
	Db       *gorm.DB
	Err 		 error
)