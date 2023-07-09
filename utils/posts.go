/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package utils

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
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
