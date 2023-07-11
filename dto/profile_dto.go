/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package dto

import "time"

type ProfileDTO struct {
	Name      string    `json:"name" validate:"lte=100"`
	Birthday  time.Time `json:"birthday"`
	Bio       string    `json:"bio" validate:"lte=240"`
	Link      string    `json:"link"`
	Image     string    `json:"image"`
	IsPrivate *bool     `json:"is_private"`
}
