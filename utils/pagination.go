/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const ItemsPerPage = 30

func GetOffset(c *fiber.Ctx) int {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}

	return (page - 1) * ItemsPerPage
}
