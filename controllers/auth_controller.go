/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import (
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SignUpDTO struct {
	Email    string    `json:"email" validate:"required,email,lte=255"`
	Password string    `json:"password" validate:"required,gte=8,lte=255"`
	Name     string    `json:"name" validate:"required,lte=255"`
	Username string    `json:"username" validate:"required,lte=255"`
	Birthday time.Time `json:"birthday"`
}

type SignInDTO struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

func SignUp(c *fiber.Ctx) error {
	signUp := new(SignUpDTO)
	if err := c.BodyParser(signUp); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": utils.ValidatorErrors(err)})
	}

	var u models.User

	result := database.DB.First(&u, "username = ?", signUp.Username)
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Username unavailable."})
	}

	result = database.DB.First(&u, "email = ?", signUp.Email)
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "An account exists with this email."})
	}

	user := models.User{Email: signUp.Email, Password: utils.HashMake(signUp.Password), Name: signUp.Name, Username: signUp.Username, Birthday: signUp.Birthday}
	r := database.DB.Create(&user)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakeAccessToken(c, &user)
}

func SignIn(c *fiber.Ctx) error {
	signIn := new(SignInDTO)
	if err := c.BodyParser(signIn); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": utils.ValidatorErrors(err)})
	}

	var user models.User
	result := database.DB.First(&user, "email = ?", signIn.Email)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid credentials."})
	}

	if !utils.HashCheck(signIn.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid credentials."})
	}

	return utils.MakeAccessToken(c, &user)
}

func SignOut(c *fiber.Ctx) error {
	var accessToken models.AccessToken

	r := database.DB.First(&accessToken, "id = ?", c.Locals("token_id"))
	if r.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Token not found."})
	}

	r = database.DB.Delete(&accessToken)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Successfully signed out."})
}
