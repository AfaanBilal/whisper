/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package routes

import (
	"github.com/AfaanBilal/whisper/controllers"
	"github.com/AfaanBilal/whisper/controllers/admin"
	"github.com/AfaanBilal/whisper/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", middleware.AuthProtected(), controllers.Home)
	app.Get("/explore", middleware.AuthProtected(), controllers.Explore)
	app.Get("/search", middleware.AuthProtected(), controllers.SearchUsers)

	auth := app.Group("/auth")
	auth.Post("/sign-up", controllers.SignUp)
	auth.Post("/sign-in", controllers.SignIn)
	auth.Post("/request-reset-password", controllers.RequestResetPassword)
	auth.Post("/verify-code", controllers.VerifyCode)
	auth.Post("/reset-password", controllers.ResetPassword)
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

	admin_routes := app.Group("/admin", middleware.AuthProtected(), middleware.AdminOnly())
	admin_routes.Get("/", admin.GetDashboard)

	admin_posts := admin_routes.Group("/posts")
	admin_posts.Get("/", admin.GetPosts)
	admin_posts.Delete("/:uuid", admin.DeletePost)

	admin_users := admin_routes.Group("/users")
	admin_users.Get("/", admin.GetUsers)
	admin_users.Delete("/:uuid", admin.DeleteUser)
}
