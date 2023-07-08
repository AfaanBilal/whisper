/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package utils

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
)

func AuthUser(c *fiber.Ctx) models.User {
	var user models.User
	result := database.DB.First(&user, "id = ?", c.Locals("user_id"))

	if result.RowsAffected == 0 {
		panic("No authenticated user found")
	}

	return user
}
