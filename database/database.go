package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/RafaelMoreira96/secure-campus-backend/database/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateUserAdmin(db *gorm.DB) {
	/*
		var user models.VoluntarioMusical
		user.NomeUsuario = "admin"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		user.SenhaUsuario = string(hashedPassword)
		user.TipoUsuario = "GESTOR_TI"
		db.AutoMigrate(models.VoluntarioMusical{})

		db.Create(&user)
	*/
}

func Connect() *gorm.DB {
	dbPath := "files"
	dbFile := filepath.Join(dbPath, "database.sqlite")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Erro ao criar diretório: %v", err)
		}
	}

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			log.Fatalf("Erro ao criar arquivo: %v", err)
		}
		file.Close()
	}

	database, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro: ", err)
		return nil
	}
	db = database
	migrations.RunMigrations(db)
	CreateUserAdmin(db)
	return db.Begin()
}

func GetDatabase() *gorm.DB {
	return db
}
