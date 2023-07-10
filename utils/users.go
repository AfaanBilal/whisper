/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
)

type UserResource struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Image     string `json:"image"`
	IsPrivate bool   `json:"is_private"`
}

func ProcessUsersResponse(c *fiber.Ctx, users []models.User) []UserResource {
	var us []UserResource
	for _, u := range users {
		us = append(us, UserResource{
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
