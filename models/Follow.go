/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package models

import (
	"time"
)

type Follow struct {
	ID uint `gorm:"primaryKey,autoIncrement"`

	FollowedId uint      `gorm:"index"`
	FollowerId uint      `gorm:"index"`
	AcceptedAt time.Time `gorm:"index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
