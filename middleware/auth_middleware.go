/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package middleware

import (
	"strings"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func AuthProtected() func(*fiber.Ctx) error {
	config := keyauth.Config{
		Validator:    authValidator,
		ErrorHandler: authError,
	}

	return keyauth.New(config)
}

func authValidator(c *fiber.Ctx, key string) (bool, error) {
	keyParts := strings.Split(key, "|")
	accessTokenId := keyParts[0]
	accessTokenValue := keyParts[1]

	var accessToken models.AccessToken
	result := database.DB.First(&accessToken, "id = ?", accessTokenId)

	if result.RowsAffected == 0 {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}

	if !utils.HashCheck(accessTokenValue, accessToken.Token) {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}

	c.Locals("token_id", accessToken.ID)
	c.Locals("user_id", accessToken.UserId)

	return true, nil
}

func authError(c *fiber.Ctx, err error) error {
	if err.Error() == keyauth.ErrMissingOrMalformedAPIKey.Error() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
