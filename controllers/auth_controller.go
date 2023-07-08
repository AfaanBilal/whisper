/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
	Name     string `json:"name" validate:"required,lte=255"`
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

	var u models.User
	result := database.DB.First(&u, "email = ?", signUp.Email)

	if result.RowsAffected > 0 {
		return c.JSON(fiber.Map{"status": "error", "message": "An account exists with this email."})
	}

	user := models.User{Email: signUp.Email, Password: utils.HashMake(signUp.Password), Name: signUp.Name}
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

	var user models.User
	result := database.DB.First(&user, "email = ?", signIn.Email)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid credentials."})
	}

	if !utils.HashCheck(signIn.Password, user.Password) {
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid credentials."})
	}

	return utils.MakeAccessToken(c, &user)
}
