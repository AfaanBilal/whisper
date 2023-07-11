/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package resources

import "time"

type ProfileResource struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Link      string     `json:"link"`
	Bio       string     `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
	Image     string     `json:"image"`
	IsPrivate bool       `json:"is_private"`
}
