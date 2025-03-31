package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// Function to hash a password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// Function to compare a hashed password with the original password
func ComparePasswords(hashedPassword string, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidatePassword(password string) error {

	digitRegex := regexp.MustCompile(`[0-9]`)
	if !digitRegex.MatchString(password) {
		return fmt.Errorf("Password did not meet the required strength. Must contain at least digit")
	}

	letterRegex := regexp.MustCompile(`[a-zA-Z]`)
	if !letterRegex.MatchString(password) {
		return fmt.Errorf("Password did not meet the required strength. Must contain at least letter")
	}

	if len(password) < 8 || len(password) > 32 {
		return fmt.Errorf("Password did not meet the required strength. Must have be 8-32 characters long")
	}

	return nil

}
