/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package controllers

import (
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/AfaanBilal/whisper/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthorResource struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

type PostResource struct {
	UUID      string         `json:"uuid"`
	Author    AuthorResource `json:"author"`
	Content   string         `json:"content"`
	Media     string         `json:"media"`
	CreatedAt time.Time      `json:"created_at"`
	Likes     uint           `json:"likes"`
	Liked     bool           `json:"liked"`
}

func Home(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Where("user_id IN (SELECT followed_id FROM follows WHERE follower_id = ?)", utils.AuthId(c)).Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	var post_ids []uint
	var user_ids []uint
	for _, post := range posts {
		user_ids = append(user_ids, post.UserId)
		post_ids = append(post_ids, post.ID)
	}

	var authors []models.User
	r = database.DB.Where("id IN ?", user_ids).Limit(30).Find(&authors)
	if r.Error != nil {
		panic(r.Error)
	}

	allLikes := utils.LikeCounts(post_ids)
	likedPosts := utils.LikedPosts(utils.AuthId(c), post_ids)

	var ps []PostResource
	for _, post := range posts {
		author := utils.FindUser(authors, post.UserId)

		ps = append(ps, PostResource{
			UUID: post.UUID.String(),
			Author: AuthorResource{
				Name:     author.Name,
				Username: author.Username,
				Image:    author.Image,
			},
			Content:   post.Content,
			Media:     post.Media,
			CreatedAt: post.CreatedAt,
			Likes:     uint(allLikes[post.ID]),
			Liked:     utils.HasLiked(utils.AuthId(c), post.ID, likedPosts),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "posts": ps})
}

func Explore(c *fiber.Ctx) error {
	var posts []models.Post

	r := database.DB.Order("id DESC").Limit(30).Find(&posts)
	if r.Error != nil {
		panic(r.Error)
	}

	return c.JSON(fiber.Map{"status": "success"})
}
