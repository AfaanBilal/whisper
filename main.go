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

	"github.com/AfaanBilal/whisper/controllers"
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/middleware"
	"github.com/AfaanBilal/whisper/models"
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

	app.Get("/", middleware.AuthProtected(), controllers.Home)
	app.Get("/explore", middleware.AuthProtected(), controllers.Explore)

	auth := app.Group("/auth")
	auth.Post("/sign-up", controllers.SignUp)
	auth.Post("/sign-in", controllers.SignIn)
	auth.Post("/sign-out", middleware.AuthProtected(), controllers.SignOut)

	me := app.Group("/me", middleware.AuthProtected())
	me.Get("/", controllers.GetProfile)
	me.Put("/", controllers.UpdateProfile)
	me.Get("/followers", controllers.GetFollowers)
	me.Get("/following", controllers.GetFollowing)
	me.Get("/notifications", controllers.GetNotifications)
	me.Post("/followers/:uuid/accept", controllers.AcceptFollower)
	me.Post("/followers/:uuid/reject", controllers.RejectFollower)
	me.Delete("/followers/:uuid", controllers.RemoveFollower)

	users := app.Group("/users", middleware.AuthProtected())
	users.Get("/:uuid", controllers.GetUserProfile)
	users.Get("/:uuid/followers", controllers.GetUserFollowers)
	users.Get("/:uuid/following", controllers.GetUserFollowing)
	users.Post("/:uuid/follow", controllers.FollowUser)
	users.Post("/:uuid/follow/cancel", controllers.CancelFollowRequest)
	users.Delete("/:uuid/follow", controllers.UnfollowUser)

	posts := app.Group("/posts", middleware.AuthProtected())
	posts.Get("/", controllers.GetPosts)
	posts.Post("/", controllers.CreatePost)
	posts.Get("/:uuid", controllers.GetPost)
	posts.Put("/:uuid", controllers.UpdatePost)
	posts.Delete("/:uuid", controllers.DeletePost)
	posts.Get("/:uuid/likes", controllers.GetLikes)
	posts.Post("/:uuid/likes", controllers.LikePost)
	posts.Delete("/:uuid/likes", controllers.UnlikePost)

	fmt.Println(fmt.Sprintf("[whisper %s] Starting on port %s.", os.Getenv("VERSION"), os.Getenv("PORT")))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
