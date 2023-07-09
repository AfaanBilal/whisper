/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import (
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	posts := utils.UserPosts(utils.AuthId(c))
	postCount := utils.PostCount(utils.AuthId(c))
	followerCount := utils.FollowerCount(utils.AuthId(c))
	followingCount := utils.FollowerCount(utils.AuthId(c))

	return c.JSON(fiber.Map{"status": "success", "profile": utils.AuthUser(c), "posts": posts, "post_count": postCount, "follower_count": followerCount, "following_count": followingCount})
}

type ProfileDTO struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Bio      string    `json:"bio"`
	Link     string    `json:"link"`
}

func UpdateProfile(c *fiber.Ctx) error {
	profileData := new(ProfileDTO)
	if err := c.BodyParser(profileData); err != nil {
		return err
	}

	user := utils.AuthUser(c)
	r := database.DB.Model(&user).Updates(profileData)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "profile": user})
}

func GetFollowers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "followers": utils.UserFollowers(utils.AuthId(c))})
}

func GetFollowing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "following": utils.UserFollowing(utils.AuthId(c))})
}

func GetNotifications(c *fiber.Ctx) error {
	var notifications []models.Notification

	r := database.DB.Where("user_id  = ?", utils.AuthId(c)).Find(&notifications)
	if r.Error != nil {
		panic("Can't find notifications")
	}

	return c.JSON(fiber.Map{"status": "success", "notifications": notifications})
}
