/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package validation

import (
	"github.com/AfaanBilal/whisper/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateProfile(c *fiber.Ctx) (*dto.ProfileDTO, error) {
	profileData := new(dto.ProfileDTO)
	if err := c.BodyParser(profileData); err != nil {
		return profileData, err
	}

	validate := validator.New()
	if err := validate.Struct(profileData); err != nil {
		return profileData, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return profileData, nil
}
