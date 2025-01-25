package db

import (
	"go-crud-app/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "app.db" // Caminho padrão se a variável de ambiente não estiver definida
	}
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrar o esquema do banco de dados
	DB.AutoMigrate(&models.Item{})
}
