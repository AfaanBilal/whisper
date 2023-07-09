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

type Notification struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	UserId uint `gorm:"index"`

	TargetPostId   uint `gorm:"index"`
	TargetUserId   uint `gorm:"index"`
	TargetFollowId uint `gorm:"index"`

	Type    string `gorm:"type:varchar(255); index"`
	Message string `gorm:"type:varchar(255)"`
	Link    string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
