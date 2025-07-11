package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"api2/src/entities"
)

var DB *gorm.DB

func Connect() {
	//dsn := "root:root@tcp(54.226.109.12:3306)/mydb?parseTime=true"
	dsn := "root:root@tcp(localhost:3307)/mydb?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	db.AutoMigrate(&entities.Visitas{}, &entities.Atraccion{}, &entities.VisitaGeneral{})
	DB = db
}
