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
	"github.com/AfaanBilal/whisper/resources"
	"github.com/gofiber/fiber/v2"
)

func ProcessNotificationResponse(c *fiber.Ctx, notifications []models.Notification) []resources.NotificationResource {
	var notification_ids []uint
	var user_ids []uint
	var follow_ids []uint
	for _, n := range notifications {
		user_ids = append(user_ids, n.TargetUserId)
		notification_ids = append(notification_ids, n.ID)
		follow_ids = append(follow_ids, n.TargetFollowId)
	}

	var users []models.User
	r := database.DB.Where("id IN ?", user_ids).Limit(30).Find(&users)
	if r.Error != nil {
		panic(r.Error)
	}

	var follows []models.Follow
	r = database.DB.Where("id IN ?", follow_ids).Find(&follows)
	if r.Error != nil {
		panic(r.Error)
	}

	var ns []resources.NotificationResource
	for _, n := range notifications {
		user := FindUser(users, n.TargetUserId)

		ns = append(ns, resources.NotificationResource{
			ID: n.ID,
			User: resources.UserResource{
				UUID:      user.UUID.String(),
				Name:      user.Name,
				Username:  user.Username,
				Image:     user.Image,
				IsPrivate: user.IsPrivate,
			},
			FollowId:  FindFollowId(follows, n.TargetFollowId),
			Type:      n.Type,
			Message:   n.Message,
			CreatedAt: n.CreatedAt,
		})
	}

	return ns
}
