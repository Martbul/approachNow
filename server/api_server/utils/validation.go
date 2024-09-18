package utils

import "regexp"

const (
	emailPattern   = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordMinLen = 8
)


func ValidateEmail(email string) bool {
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

func ValidatePassword(password string) bool {
	return len(password) >= passwordMinLen
}
