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

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id =?", utils.AuthId(c)).Find(&posts)
	if r.Error != nil {
		panic("Can't find posts")
	}

	return c.JSON(fiber.Map{"status": "success", "posts": posts})
}

type PostDTO struct {
	Content string `json:"content"`
	Media   string `json:"media"`
}

func CreatePost(c *fiber.Ctx) error {
	postData := new(PostDTO)
	if err := c.BodyParser(postData); err != nil {
		return err
	}

	post := models.Post{UserId: utils.AuthId(c), Content: postData.Content, Media: postData.Media}
	r := database.DB.Create(&post)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "post": post})
}

func UpdatePost(c *fiber.Ctx) error {
	postData := new(PostDTO)
	if err := c.BodyParser(postData); err != nil {
		return err
	}

	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	if post.UserId != utils.AuthId(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized."})
	}

	r := database.DB.Model(&post).Updates(postData)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "post": post})
}

func GetPost(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	return c.JSON(fiber.Map{"status": "success", "post": post})
}

func DeletePost(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	if post.UserId != utils.AuthId(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized."})
	}

	r := database.DB.Delete(&post)
	if r.Error != nil {
		panic(r.Error)
	}

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
