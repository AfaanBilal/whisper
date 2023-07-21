/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package controllers

import (
	"fmt"
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/AfaanBilal/whisper/validation"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	signUp, err := validation.ValidateSignUp(c)
	if err != nil {
		return err
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

	user := models.User{Email: signUp.Email, Password: utils.HashMake(signUp.Password), Name: signUp.Name, Username: signUp.Username}

	r := database.DB.Create(&user)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakeAccessToken(c, &user)
}

func SignIn(c *fiber.Ctx) error {
	signIn, err := validation.ValidateSignIn(c)
	if err != nil {
		return err
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

func RequestResetPassword(c *fiber.Ctx) error {
	resetPassword, err := validation.ValidateRequestResetPassword(c)
	if err != nil {
		return err
	}

	var user models.User
	result := database.DB.First(&user, "email = ?", resetPassword.Email)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No account associated with that email."})
	}

	vc := models.VerificationCode{
		UserId:    user.ID,
		Code:      utils.MakeCode(),
		Token:     utils.MakeToken(),
		Purpose:   models.Purpose_PasswordReset,
		Attempts:  3,
		ExpiresAt: time.Now().Add(time.Minute * 15),
	}

	r := database.DB.Create(&vc)
	if r.Error != nil {
		panic(r.Error)
	}

	utils.SendPasswordResetCode(user.Email, user.Name, vc.Code)

	return c.JSON(fiber.Map{"status": "success", "message": "A verification code has been sent to your email.", "uuid": vc.UUID})
}

func VerifyCode(c *fiber.Ctx) error {
	verifyCode, err := validation.ValidateVerifyCode(c)
	if err != nil {
		return err
	}

	var vc models.VerificationCode
	result := database.DB.First(&vc, "uuid = ?", verifyCode.UUID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid VC UUID."})
	}

	if vc.Attempts <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Maximum attempts exceeded. Please request a new code."})
	}

	if vc.ExpiresAt.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Code expired. Please request a new code."})
	}

	if vc.Code != verifyCode.Code {
		vc.Attempts -= 1
		database.DB.Updates(&vc)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid code. Attempts remaining: " + fmt.Sprintf("%d", vc.Attempts)})
	}

	vc.Attempts = 0
	database.DB.Updates(&vc)

	return c.JSON(fiber.Map{"status": "success", "uuid": vc.UUID, "token": vc.Token})
}

func ResetPassword(c *fiber.Ctx) error {
	resetPassword, err := validation.ValidateResetPassword(c)
	if err != nil {
		return err
	}

	var vc models.VerificationCode
	result := database.DB.First(&vc, "uuid = ?", resetPassword.UUID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid VC UUID."})
	}

	if vc.Token != resetPassword.Token {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Code invalid or expired. Please request a new code."})
	}

	var user models.User
	result = database.DB.First(&user, "id = ?", vc.UserId)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid VC user."})
	}

	user.Password = utils.HashMake(resetPassword.Password)
	result = database.DB.Model(&user).Update("password", user.Password)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something went wrong. Please try again. E_PW_UPDATE_FAILED."})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Password reset complete."})
}
