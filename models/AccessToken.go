/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package models

import (
	"time"
)

type AccessToken struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	UserId uint `gorm:"index"`

	Name  string `gorm:"type:varchar(255)"`
	Token string `gorm:"type:varchar(255); uniqueIndex"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
