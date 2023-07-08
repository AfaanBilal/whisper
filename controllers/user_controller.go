/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	user, err := models.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	var posts []models.Post
	r := database.DB.Where("user_id = ?", user.ID).Limit(20).Find(&posts)
	if r.Error != nil {
		panic("Can't find posts")
	}

	return c.JSON(fiber.Map{"status": "success", "profile": user, "posts": posts})
}

func GetUserFollowers(c *fiber.Ctx) error {
	user, err := models.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	followers := utils.UserFollowers(user.ID)

	return c.JSON(fiber.Map{"status": "success", "followers": followers})
}

func GetUserFollowing(c *fiber.Ctx) error {
	user, err := models.GetUser(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	following := utils.UserFollowing(user.ID)

	return c.JSON(fiber.Map{"status": "success", "following": following})
}
