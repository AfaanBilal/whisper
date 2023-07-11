/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package resources

import "time"

type NotificationResource struct {
	ID        uint         `json:"id"`
	User      UserResource `json:"user"`
	FollowId  uint         `json:"follow_id"`
	Type      string       `json:"type"`
	Message   string       `json:"message"`
	CreatedAt time.Time    `json:"created_at"`
}
