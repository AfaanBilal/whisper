/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/resources"
	"github.com/gofiber/fiber/v2"
)

func GetUser(uuid string) (models.User, error) {
	var user models.User

	result := database.DB.First(&user, "uuid = ?", uuid)
	if result.RowsAffected == 0 {
		return user, result.Error
	}

	return user, nil
}

func ProcessUsersResponse(c *fiber.Ctx, users []models.User) []resources.UserResource {
	var us []resources.UserResource
	for _, u := range users {
		us = append(us, resources.UserResource{
			UUID:      u.UUID.String(),
			Name:      u.Name,
			Username:  u.Username,
			Image:     u.Image,
			IsPrivate: u.IsPrivate,
		})
	}

	return us
}

func MakeUsersResponse(c *fiber.Ctx, users []models.User) error {
	return c.JSON(fiber.Map{"status": "success", "users": ProcessUsersResponse(c, users)})
}

func TotalUserCount() int64 {
	var user models.User
	var count int64
	database.DB.Model(&user).Count(&count)
	return count
}
