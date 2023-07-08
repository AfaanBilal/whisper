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
