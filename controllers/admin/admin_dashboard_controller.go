/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package admin

import (
	"github.com/gofiber/fiber/v2"
)

func GetDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success"})
}
