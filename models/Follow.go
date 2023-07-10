/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package models

import (
	"database/sql"
	"time"
)

type Follow struct {
	ID uint `gorm:"primaryKey,autoIncrement"`

	FollowedId uint         `gorm:"index"`
	FollowerId uint         `gorm:"index"`
	AcceptedAt sql.NullTime `gorm:"index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
