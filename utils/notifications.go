/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"time"

	"github.com/AfaanBilal/whisper/database"
	"github.com/AfaanBilal/whisper/models"
	"github.com/gofiber/fiber/v2"
)

type NotificationResource struct {
	ID        uint         `json:"id"`
	User      UserResource `json:"user"`
	FollowId  uint         `json:"follow_id"`
	Type      string       `json:"type"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at"`
}

func ProcessNotificationResponse(c *fiber.Ctx, notifications []models.Notification) []NotificationResource {
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

	var ns []NotificationResource
	for _, n := range notifications {
		user := FindUser(users, n.TargetUserId)

		ns = append(ns, NotificationResource{
			ID: n.ID,
			User: UserResource{
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
