/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"fmt"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func MakeAccessToken(c *fiber.Ctx, user *models.User) error {
	token := uuid.New().String()
	accessToken := models.AccessToken{UserId: user.ID, Name: "login", Token: HashMake(token)}
	r := database.DB.Create(&accessToken)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "access_token": fmt.Sprintf("%d", accessToken.ID) + "|" + token})
}
