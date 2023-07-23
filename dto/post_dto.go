/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package dto

type PostDTO struct {
	Content string `json:"content" validate:"lte=233"`
	Media   string `json:"media"`
}
