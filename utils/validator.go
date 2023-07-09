/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidatorErrors(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		return err.Field() + " is invalid."
	}

	return ""
}
