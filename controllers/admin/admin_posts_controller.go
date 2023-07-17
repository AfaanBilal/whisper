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

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Order("id DESC").Offset(utils.GetOffset(c)).Limit(utils.ItemsPerPage).Find(&posts)
	if r.Error != nil {
		panic("Can't find posts")
	}

	return c.JSON(fiber.Map{"status": "success", "posts": utils.ProcessPostsResponse(c, posts)})
}

func DeletePost(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	r := database.DB.Delete(&post)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}
