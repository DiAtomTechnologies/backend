package models

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type User struct {
	Id       string `json:"id,omitempty"`
    Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

func (u *User) ValidateUser() error {
	if !isValidEmail(u.Email) {
      return errors.New("Invalid email provider. we only accept mail from the following providers: gmail, yahoo and outlook")
	}

	if !isValidPassword(u.Password) {
		return errors.New("Password must atleast have more than 8 characters and contain atleast one upper , one lower and one digit.")
	}
	role := strings.ToUpper(u.Role)
	if role != "ADMIN" && role != "MANAGER" {
		return errors.New("invalid role")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@(gmail|outlook|yahoo)+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
func isValidPassword(password string) bool {
	var hasUpper, hasLower, hasDigit bool
	if len(password) >= 8 {
		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsDigit(char):
				hasDigit = true
			}
		}
	}
	return hasUpper && hasLower && hasDigit
}
