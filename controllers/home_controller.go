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
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id IN (SELECT followed_id FROM follows WHERE follower_id = ? AND accepted_at IS NOT NULL)", utils.AuthId(c)).Or("user_id = ?", utils.AuthId(c)).Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakePostsResponse(c, posts)
}

func Explore(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id NOT IN (SELECT id FROM users WHERE is_private = 1)").Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakePostsResponse(c, posts)
}

func SearchUsers(c *fiber.Ctx) error {
	search := c.Query("s", "")
	search = "%" + search + "%"

	var users []models.User
	if search == "" {
		return utils.MakeUsersResponse(c, users)
	}

	r := database.DB.Where("name LIKE ? OR username LIKE ?", search, search).Limit(20).Find(&users)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakeUsersResponse(c, users)
}
