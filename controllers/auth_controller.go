/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import "github.com/gofiber/fiber/v2"

func SignUp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func SignIn(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}
