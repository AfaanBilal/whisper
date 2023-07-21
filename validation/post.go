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

func ValidatePost(c *fiber.Ctx) (*dto.PostDTO, error) {
	postData := new(dto.PostDTO)
	if err := c.BodyParser(postData); err != nil {
		return postData, err
	}

	validate := validator.New()
	if err := validate.Struct(postData); err != nil {
		return postData, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return postData, nil
}
