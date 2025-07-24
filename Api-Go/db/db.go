package db

import (
	"log"
	//Para el railway
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global para que no este eclarado en cada handlers
var DB *gorm.DB

func Connect() {
	// Aqui se pilla da la databse url que ofrece railway, en cado contrario usa la config local
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Si no hay DATABASE_URL, usamos configuración local por defecto
		dsn = "host=localhost user=postgres password=root dbname=SupportApp port=5432 sslmode=disable TimeZone=Europe/Madrid"
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	DB = connection
}
