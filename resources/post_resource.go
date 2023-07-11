/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package resources

import "time"

type PostResource struct {
	UUID      string       `json:"uuid"`
	Author    UserResource `json:"author"`
	Content   string       `json:"content"`
	Media     string       `json:"media"`
	CreatedAt time.Time    `json:"created_at"`
	Likes     uint         `json:"likes"`
	Liked     bool         `json:"liked"`
}
