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

func ProcessProfileResponse(user models.User) resources.ProfileResource {
	p := resources.ProfileResource{
		UUID:      user.UUID.String(),
		Name:      user.Name,
		Username:  user.Username,
		Link:      user.Link,
		Bio:       user.Bio,
		Birthday:  nil,
		Image:     user.Image,
		IsPrivate: user.IsPrivate,
	}

	if user.Birthday.Valid {
		p.Birthday = &user.Birthday.Time
	}

	return p
}

func AuthUser(c *fiber.Ctx) models.User {
	var user models.User
	result := database.DB.First(&user, "id = ?", c.Locals("user_id"))

	if result.RowsAffected == 0 {
		panic("No authenticated user found")
	}

	return user
}

func AuthId(c *fiber.Ctx) uint {
	return c.Locals("user_id").(uint)
}
