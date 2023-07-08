/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package main

import (
	"fmt"
	"os"

	"github.com/AfaanBilal/whisper/controllers"
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.AccessToken{}, &models.Post{}, &models.Like{})

	app := fiber.New(fiber.Config{
		ServerHeader:          "Whisper",
		AppName:               "Whisper " + os.Getenv("VERSION"),
		DisableStartupMessage: true,
	})

	app.Use(cors.New())

	app.Get("/", controllers.Home)
	app.Get("/explore", controllers.Explore)

	auth := app.Group("/auth")
	auth.Post("/sign-up", controllers.SignUp)
	auth.Post("/sign-in", controllers.SignIn)

	me := app.Group("/me")
	me.Get("/", controllers.GetProfile)
	me.Put("/", controllers.UpdateProfile)
	me.Get("/followers", controllers.GetFollowers)
	me.Get("/following", controllers.GetFollowing)

	posts := app.Group("/posts")
	posts.Get("/", controllers.GetPosts)
	posts.Post("/", controllers.CreatePost)
	posts.Get("/:uuid", controllers.GetPost)
	posts.Put("/:uuid", controllers.UpdatePost)
	posts.Delete("/:uuid", controllers.DeletePost)
	posts.Get("/:uuid/likes", controllers.LikePost)
	posts.Post("/:uuid/likes", controllers.LikePost)
	posts.Delete("/:uuid/likes", controllers.UnlikePost)
	posts.Post("/:uuid/reply", controllers.ReplyPost)

	fmt.Println(fmt.Sprintf("[whisper %s] Starting on port %s.", os.Getenv("VERSION"), os.Getenv("PORT")))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
