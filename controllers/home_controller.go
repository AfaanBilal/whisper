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

func Home(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id IN (SELECT followed_id FROM follows WHERE follower_id = ?)", utils.AuthId(c)).Or("user_id = ?", utils.AuthId(c)).Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakePostsResponse(c, posts)
}

func Explore(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	return utils.MakePostsResponse(c, posts)
}
