/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint          `gorm:"primaryKey,autoIncrement" json:"-"`
	UUID      uuid.UUID     `gorm:"type:varchar(60);" json:"uuid"`
	UserId    uint          `json:"-"`
	ReplyToId sql.NullInt64 `json:"-"`

	Content string          `json:"content"`
	Media   *sql.NullString `json:"media"`

	Meta      *string   `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.UUID = uuid.New()
	return
}
