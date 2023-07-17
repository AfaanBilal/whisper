/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package middleware

import (
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func AdminOnly() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := utils.AuthUser(c)
		if user.Role != "admin" {
			return c.SendStatus(fiber.StatusForbidden)
		}

		return c.Next()
	}
}
