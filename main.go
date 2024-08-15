package main

import (
	"fmt"
	"log"

	"github.com/ggorockee/airbnb-backend/config"
	"github.com/ggorockee/airbnb-backend/controller"
	"github.com/ggorockee/airbnb-backend/model"
	"github.com/ggorockee/airbnb-backend/repository"
	"github.com/ggorockee/airbnb-backend/router"
	"github.com/ggorockee/airbnb-backend/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Run service ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Ini Controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
