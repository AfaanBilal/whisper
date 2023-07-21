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

func ValidateSignUp(c *fiber.Ctx) (*dto.SignUpDTO, error) {
	signUp := new(dto.SignUpDTO)
	if err := c.BodyParser(signUp); err != nil {
		return signUp, err
	}

	validate := validator.New()
	if err := validate.Struct(signUp); err != nil {
		return signUp, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return signUp, nil
}

func ValidateSignIn(c *fiber.Ctx) (*dto.SignInDTO, error) {
	signIn := new(dto.SignInDTO)
	if err := c.BodyParser(signIn); err != nil {
		return signIn, err
	}

	validate := validator.New()
	if err := validate.Struct(signIn); err != nil {
		return signIn, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return signIn, nil
}

func ValidateRequestResetPassword(c *fiber.Ctx) (*dto.RequestResetPasswordDTO, error) {
	resetPassword := new(dto.RequestResetPasswordDTO)
	if err := c.BodyParser(resetPassword); err != nil {
		return resetPassword, err
	}

	validate := validator.New()
	if err := validate.Struct(resetPassword); err != nil {
		return resetPassword, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return resetPassword, nil
}

func ValidateVerifyCode(c *fiber.Ctx) (*dto.VerifyCodeDTO, error) {
	verifyCode := new(dto.VerifyCodeDTO)
	if err := c.BodyParser(verifyCode); err != nil {
		return verifyCode, err
	}

	validate := validator.New()
	if err := validate.Struct(verifyCode); err != nil {
		return verifyCode, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return verifyCode, nil
}

func ValidateResetPassword(c *fiber.Ctx) (*dto.ResetPasswordDTO, error) {
	resetPassword := new(dto.ResetPasswordDTO)
	if err := c.BodyParser(resetPassword); err != nil {
		return resetPassword, err
	}

	validate := validator.New()
	if err := validate.Struct(resetPassword); err != nil {
		return resetPassword, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": ValidatorErrors(err)})
	}

	return resetPassword, nil
}
