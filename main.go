/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package main

import (
	"fmt"
	"os"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.AccessToken{}, &models.Follow{}, &models.Post{}, &models.Like{}, &models.Notification{})

	app := fiber.New(fiber.Config{
		ServerHeader:          "Whisper",
		AppName:               "Whisper " + os.Getenv("VERSION"),
		DisableStartupMessage: true,
	})

	app.Use(cors.New())

	routes.Setup(app)

	fmt.Println(fmt.Sprintf("[whisper %s] Starting on port %s.", os.Getenv("VERSION"), os.Getenv("PORT")))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
