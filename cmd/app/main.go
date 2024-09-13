package main

import (
	"br_api/config"
	"br_api/db"
	"br_api/internal/models/entity"
	"br_api/internal/repository"
	"br_api/internal/service"
	"br_api/internal/transport"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	repo := initializeRepository(cfg)
	braineeService := service.NewBraineeService(repo)
	braineeHandler := transport.NewBraineeHandler(braineeService)
	router := transport.NewRouter(braineeHandler)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}

func initializeRepository(cfg *config.Config) repository.BraineeRepository {
	database, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Println("Error connecting to database, using in-memory storage")
		return repository.NewBraineeMemoryRepository()
	}

	if err := database.AutoMigrate(&entity.Brainee{}); err != nil {
		log.Println("Error migrating database, using in-memory storage")
		return repository.NewBraineeMemoryRepository()
	}

	return repository.NewBraineeDBRepository(database)
}
