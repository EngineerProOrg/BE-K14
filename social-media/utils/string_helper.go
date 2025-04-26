package utils

import "strings"

func GetUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	var username string
	if len(parts) > 0 {
		username = parts[0]
	}
	return username
}
