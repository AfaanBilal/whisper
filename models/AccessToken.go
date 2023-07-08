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

type AccessToken struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	UserId uint

	Name  string `gorm:"type:varchar(255)"`
	Token string

	CreatedAt time.Time
	UpdatedAt time.Time
}