/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidatorErrors(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Printf("%s", err.Error())
		return err.Field() + " is invalid."
	}

	return ""
}
