/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
)

func UserFollowers(userId uint) []models.User {
	var follows []models.Follow

	r := database.DB.Where("followed_id = ? AND accepted_at IS NOT NULL", userId).Find(&follows)
	if r.Error != nil {
		panic("Can't find followers")
	}

	var follower_ids []uint
	for _, follow := range follows {
		follower_ids = append(follower_ids, follow.FollowerId)
	}

	var followers []models.User
	r = database.DB.Where("id IN ?", follower_ids).Find(&followers)
	if r.Error != nil {
		panic(r.Error)
	}

	return followers
}

func UserFollowing(userId uint) []models.User {
	var follows []models.Follow

	r := database.DB.Where("follower_id = ? AND accepted_at IS NOT NULL", userId).Find(&follows)
	if r.Error != nil {
		panic("Can't find following")
	}

	var following_ids []uint
	for _, follow := range follows {
		following_ids = append(following_ids, follow.FollowedId)
	}

	var following []models.User
	r = database.DB.Where("id IN ?", following_ids).Find(&following)
	if r.Error != nil {
		panic(r.Error)
	}

	return following
}

func FollowerCount(userId uint) int64 {
	var follow models.Follow
	var count int64
	database.DB.Where("followed_id = ? AND accepted_at IS NOT NULL", userId).Model(&follow).Count(&count)
	return count
}

func FollowingCount(userId uint) int64 {
	var follow models.Follow
	var count int64
	database.DB.Where("follower_id = ? AND accepted_at IS NOT NULL", userId).Model(&follow).Count(&count)
	return count
}

func IsFollowed(userId uint, by uint) bool {
	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", userId, by)
	return result.RowsAffected > 0 && follow.AcceptedAt.Valid
}

func IsFollower(userId uint, by uint) bool {
	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", by, userId)
	return result.RowsAffected > 0 && follow.AcceptedAt.Valid
}

func IsFollowRequested(userId uint, by uint) bool {
	var follow models.Follow
	result := database.DB.First(&follow, "followed_id = ? AND follower_id = ?", userId, by)
	return result.RowsAffected > 0 && !follow.AcceptedAt.Valid
}
