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
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	user := utils.AuthUser(c)
	return c.JSON(fiber.Map{"status": "success", "profile": user})
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
	return c.JSON(fiber.Map{"status": "success"})
}

func GetFollowing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}
