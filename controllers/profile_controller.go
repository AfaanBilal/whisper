/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import "github.com/gofiber/fiber/v2"

func GetProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "my_user_id": c.Locals("user_id")})
}

func UpdateProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func GetFollowers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func GetFollowing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}
