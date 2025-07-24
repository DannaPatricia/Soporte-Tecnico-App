package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global para que no este eclarado en cada handlers
var DB *gorm.DB

func Connect() {
	// Configuracion con el server de postgres cambia pass o user si ves que no va
	dsn := "host=localhost user=postgres password=root dbname=SupportApp port=5432 sslmode=disable TimeZone=Europe/Madrid"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	DB = connection
}
