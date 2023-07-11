/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package controllers

import (
	"database/sql"
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/resources"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	followed := utils.IsFollowed(user.ID, utils.AuthId(c))
	follower := utils.IsFollower(user.ID, utils.AuthId(c))

	requested := utils.IsFollowRequested(user.ID, utils.AuthId(c))

	postCount := utils.PostCount(user.ID)
	followerCount := utils.FollowerCount(user.ID)
	followingCount := utils.FollowerCount(user.ID)

	if !user.IsPrivate || followed {
		posts := utils.ProcessPostsResponse(c, utils.UserPosts(user.ID))

		return c.JSON(fiber.Map{
			"status":          "success",
			"profile":         user,
			"followed":        followed,
			"follower":        follower,
			"posts":           posts,
			"post_count":      postCount,
			"follower_count":  followerCount,
			"following_count": followingCount,
		})
	} else {
		return c.JSON(fiber.Map{
			"status": "success",
			"profile": resources.UserResource{
				UUID:      user.UUID.String(),
				Name:      user.Name,
				Username:  user.Username,
				Image:     user.Image,
				IsPrivate: user.IsPrivate,
			},
			"requested":       requested,
			"followed":        followed,
			"follower":        follower,
			"post_count":      postCount,
			"follower_count":  followerCount,
			"following_count": followingCount,
		})
	}
}

func GetUserFollowers(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	followers := utils.UserFollowers(user.ID)

	return c.JSON(fiber.Map{"status": "success", "followers": followers})
}

func GetUserFollowing(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	following := utils.UserFollowing(user.ID)

	return c.JSON(fiber.Map{"status": "success", "following": following})
}

func FollowUser(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Already followed."})
	}

	follow = models.Follow{FollowedId: user.ID, FollowerId: utils.AuthId(c)}
	if !user.IsPrivate {
		follow.AcceptedAt = sql.NullTime{Time: time.Now(), Valid: true}
	}

	r := database.DB.Create(&follow)
	if r.Error != nil {
		panic(r.Error)
	}

	Type := "follow"
	Message := "followed you."
	if user.IsPrivate {
		Type = "follow-requested"
		Message = "requested to follow you."
	}

	database.DB.Create(&models.Notification{UserId: user.ID, TargetUserId: utils.AuthId(c), TargetFollowId: follow.ID, Type: Type, Message: Message})

	followers := utils.UserFollowers(user.ID)

	return c.JSON(fiber.Map{"status": "success", "followers": followers})
}

func CancelFollowRequest(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Not requested."})
	}

	r := database.DB.Delete(&follow)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}

func UnfollowUser(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Not following."})
	}

	r := database.DB.Delete(&follow)
	if r.Error != nil {
		panic(r.Error)
	}

	followers := utils.UserFollowers(user.ID)

	return c.JSON(fiber.Map{"status": "success", "followers": followers})
}

func AcceptFollower(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "follower_id = ? AND followed_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Follower not found."})
	}

	r := database.DB.Model(&follow).Update("accepted_at", time.Now())
	if r.Error != nil {
		panic(r.Error)
	}

	database.DB.Create(&models.Notification{UserId: user.ID, TargetUserId: utils.AuthId(c), Type: "follow-accept", Message: "accepted your follow request."})

	return c.JSON(fiber.Map{"status": "success"})
}

func RejectFollower(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "follower_id = ? AND followed_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Follower not found."})
	}

	r := database.DB.Delete(&follow)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}

func RemoveFollower(c *fiber.Ctx) error {
	user, err := utils.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var follow models.Follow
	result := database.DB.First(&follow, "follower_id = ? AND followed_id = ?", user.ID, utils.AuthId(c))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Follower not found."})
	}

	r := database.DB.Delete(&follow)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}
