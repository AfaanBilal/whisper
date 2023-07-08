/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package utils

import (
	"fmt"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashMake(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func HashCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func MakeAccessToken(c *fiber.Ctx, user *models.User) error {
	token := uuid.New().String()
	accessToken := models.AccessToken{UserId: user.ID, Name: "login", Token: HashMake(token)}
	r := database.DB.Create(&accessToken)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "access_token": fmt.Sprintf("%d", accessToken.ID) + "|" + token})
}
