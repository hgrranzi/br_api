package main

import (
	"br_api/config"
	"br_api/db"
	"br_api/internal/models/entity"
	"br_api/internal/repository"
	"br_api/internal/service"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	var err error
	var database *gorm.DB

	database, err = db.NewPostgresDB(cfg)
	var repo repository.BraineeRepository
	if err != nil {
		log.Println("Error connecting to database, using in-memory storage")
		repo = repository.NewBraineeMemoryRepository()
	} else {
		database.AutoMigrate(entity.Brainee{})
		repo = repository.NewBraineeDBRepository(database)
	}

	braineeService := service.NewBraineeService(repo)
	braineeService.GetBraineeById(1)

}
