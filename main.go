package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zohaibsoomro/go-course-api-postgresql/db"
	"github.com/zohaibsoomro/go-course-api-postgresql/models"
)

func main() {
	godotenv.Load(".env")
	app := fiber.New()

	c := models.NewConfig(os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("SLLMODE"))
	mdb := db.NewDbConnection(c)
	mdb.SetUpRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
