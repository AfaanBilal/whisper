/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package admin

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	r := database.DB.Order("id DESC").Offset(utils.GetOffset(c)).Limit(utils.ItemsPerPage).Find(&users)
	if r.Error != nil {
		panic("Can't find users")
	}

	return c.JSON(fiber.Map{"status": "success", "users": users})
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	result := database.DB.First(&user, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found."})
	}

	r := database.DB.Delete(&user)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}
