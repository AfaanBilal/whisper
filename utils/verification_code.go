/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

func MakeCode() string {
	return RandStringCharset("0123456789", 6)
}

func MakeToken() string {
	return RandString(64)
}
