/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import "github.com/gofiber/fiber/v2"

func GetPosts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func CreatePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func UpdatePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func GetPost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func DeletePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func LikePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func UnlikePost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}

func ReplyPost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}
