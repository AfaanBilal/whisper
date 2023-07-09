/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package utils

import (
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
)

func UserPosts(userId uint) []models.Post {
	var posts []models.Post
	r := database.DB.Where("user_id = ?", userId).Limit(20).Find(&posts)
	if r.Error != nil {
		panic("Can't find posts")
	}

	return posts
}

func PostCount(userId uint) int64 {
	var post models.Post
	var count int64
	database.DB.Where("user_id = ?", userId).Model(&post).Count(&count)
	return count
}

func FindUser(users []models.User, userId uint) models.User {
	for _, user := range users {
		if user.ID == userId {
			return user
		}
	}

	return users[0]
}

func LikeCount(postId uint) int64 {
	var like models.Like
	var count int64
	database.DB.Where("post_id = ?", postId).Model(&like).Count(&count)
	return count
}

func LikeCounts(postIds []uint) map[uint]int64 {
	var like models.Like
	var counts map[uint]int64
	database.DB.Model(&like).Where("post_id IN ?", postIds).Select("post_id, COUNT(post_id) as count").Group("post_id").Find(&counts)
	return counts
}

func LikedPosts(userId uint, postIds []uint) []models.Like {
	var likes []models.Like
	database.DB.Where("user_id = ? AND post_id IN ?", userId, postIds).Find(&likes)
	return likes
}

func HasLiked(userId uint, postId uint, likes []models.Like) bool {
	for _, like := range likes {
		if like.PostId == postId && like.UserId == userId {
			return true
		}
	}

	return false
}

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

func MakePostsResponse(c *fiber.Ctx, posts []models.Post) error {
	var post_ids []uint
	var user_ids []uint
	for _, post := range posts {
		user_ids = append(user_ids, post.UserId)
		post_ids = append(post_ids, post.ID)
	}

	var authors []models.User
	r := database.DB.Where("id IN ?", user_ids).Limit(30).Find(&authors)
	if r.Error != nil {
		panic(r.Error)
	}

	allLikes := LikeCounts(post_ids)
	likedPosts := LikedPosts(AuthId(c), post_ids)

	var ps []PostResource
	for _, post := range posts {
		author := FindUser(authors, post.UserId)

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
			Liked:     HasLiked(AuthId(c), post.ID, likedPosts),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "posts": ps})
}
