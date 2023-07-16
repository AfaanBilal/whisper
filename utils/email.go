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
	"os"

	"github.com/resendlabs/resend-go"
)

func SendPasswordResetCode(email string, name string, code string) {
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Whisper <accounts@whisper.amx.gg>",
		To:      []string{email},
		Html:    fmt.Sprintf("Hi %s,<br /><br />Here's your code to reset your password: <b>%s</b><br /><br />You can ignore this message if you didn't request to reset your password.<br /><br />Team Whisper", name, code),
		Subject: "Whisper: Reset Password",
		ReplyTo: "whisper@afaan.dev",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
