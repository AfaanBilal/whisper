/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package controllers

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/AfaanBilal/whisper/validation"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id =?", utils.AuthId(c)).Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic("Can't find posts")
	}

	return c.JSON(fiber.Map{"status": "success", "posts": posts})
}

func CreatePost(c *fiber.Ctx) error {
	postData, err := validation.ValidatePost(c)
	if err != nil {
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
	postData, err := validation.ValidatePost(c)
	if err != nil {
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

func GetLikes(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	var likes []models.Like
	r := database.DB.Where("post_id =?", post.ID).Find(&likes)
	if r.Error != nil {
		panic("Can't find likes")
	}

	return c.JSON(fiber.Map{"status": "success", "likes": likes})
}

func LikePost(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	var like models.Like
	result = database.DB.First(&like, "user_id = ? AND post_id = ?", utils.AuthId(c), post.ID)
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Already liked."})
	}

	like = models.Like{UserId: utils.AuthId(c), PostId: post.ID}
	r := database.DB.Create(&like)
	if r.Error != nil {
		panic(r.Error)
	}

	database.DB.Create(&models.Notification{UserId: post.UserId, TargetUserId: utils.AuthId(c), Type: "like", Message: "liked your post."})

	return c.JSON(fiber.Map{"status": "success"})
}

func UnlikePost(c *fiber.Ctx) error {
	var post models.Post

	result := database.DB.First(&post, "uuid = ?", c.Params("uuid"))
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found."})
	}

	var like models.Like
	result = database.DB.First(&like, "user_id = ? AND post_id = ?", utils.AuthId(c), post.ID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Like not found."})
	}

	r := database.DB.Delete(&like)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}
