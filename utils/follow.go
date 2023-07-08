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

func UserFollowers(userId uint) []models.User {
	var follows []models.Follow

	r := database.DB.Where("followed_id = ?", userId).Find(&follows)
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

	r := database.DB.Where("follower_id = ?", userId).Find(&follows)
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
