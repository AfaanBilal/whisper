/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package dto

type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
	Name     string `json:"name" validate:"required,lte=255"`
	Username string `json:"username" validate:"required,lte=255"`
}

type SignInDTO struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type RequestResetPasswordDTO struct {
	Email string `json:"email" validate:"required,email,lte=255"`
}

type VerifyCodeDTO struct {
	UUID string `json:"uuid" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type ResetPasswordDTO struct {
	UUID     string `json:"uuid" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}
