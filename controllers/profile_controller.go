/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package controllers

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/dto"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/AfaanBilal/whisper/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	posts := utils.ProcessPostsResponse(c, utils.UserPosts(utils.AuthId(c)))
	postCount := utils.PostCount(utils.AuthId(c))
	followerCount := utils.FollowerCount(utils.AuthId(c))
	followingCount := utils.FollowerCount(utils.AuthId(c))

	return c.JSON(fiber.Map{
		"status":          "success",
		"profile":         utils.ProcessProfileResponse(utils.AuthUser(c)),
		"posts":           posts,
		"post_count":      postCount,
		"follower_count":  followerCount,
		"following_count": followingCount,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	profileData := new(dto.ProfileDTO)
	if err := c.BodyParser(profileData); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(profileData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": validation.ValidatorErrors(err)})
	}

	user := utils.AuthUser(c)
	r := database.DB.Model(&user).Select("*").Updates(profileData)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "profile": utils.ProcessProfileResponse(user)})
}

func GetFollowers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "followers": utils.UserFollowers(utils.AuthId(c))})
}

func GetFollowing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "following": utils.UserFollowing(utils.AuthId(c))})
}

func GetNotifications(c *fiber.Ctx) error {
	var notifications []models.Notification

	r := database.DB.Where("user_id  = ?", utils.AuthId(c)).Order("id DESC").Limit(20).Find(&notifications)
	if r.Error != nil {
		panic("Can't find notifications")
	}

	return c.JSON(fiber.Map{"status": "success", "notifications": utils.ProcessNotificationResponse(c, notifications)})
}
