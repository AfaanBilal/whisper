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

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const Purpose_SignIn = "sign-in"
const Purpose_PasswordReset = "password-reset"
const Purpose_EmailVerification = "email-verification"

type VerificationCode struct {
	ID     uint      `gorm:"primaryKey,autoIncrement"`
	UUID   uuid.UUID `gorm:"type:varchar(60); uniqueIndex" json:"uuid"`
	UserId uint      `gorm:"index"`

	Code      string `gorm:"type:varchar(50);"`
	Token     string `gorm:"type:varchar(255);"`
	Purpose   string `gorm:"type:varchar(255);"`
	Attempts  uint   `gorm:"type:tinyint;"`
	ExpiresAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (vc *VerificationCode) BeforeCreate(tx *gorm.DB) (err error) {
	vc.UUID = uuid.New()
	return
}
