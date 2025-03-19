package main

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nathaelb/authcrux/application"
	"github.com/nathaelb/authcrux/application/service"
	"github.com/nathaelb/authcrux/infrastructure/database"
	"github.com/nathaelb/authcrux/infrastructure/repository"
	"gorm.io/gorm"
)

func main() {
	isProd := os.Getenv("ENV") == "production"
	setupLogger(isProd)

	db, err := setupDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	realmRepo := repository.NewPostgresRealmRepository(db)
	realmService := service.NewRealmService(realmRepo)

	httpServer := application.NewHTTPServer(realmService)

	log.Fatal(httpServer.Start())
}

func setupLogger(isProd bool) {
	if isProd {
		log.SetLevel(log.LevelInfo)
	} else {
		log.SetLevel(log.LevelDebug)
	}
}

func setupDatabase() (*gorm.DB, error) {
	log.Info("Connecting to database...")

	dbConfig := database.NewDefaultConfig()
	db, err := database.ConnectDatabase(dbConfig)
	if err != nil {
		return nil, err
	}

	log.Info("Migrate database...")
	if err := database.MigrateDatabase(db); err != nil {
		return nil, err
	}

	return db, nil
}
