package models

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

// User represents a user in the system
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Validate validates the user data
func (u *User) Validate() error {
	log.Printf("Validating user data: %+v", u)

	if u.Name == "" {
		err := errors.New("name is required")
		log.Printf("Validation error: %v", err)
		return err
	}

	if u.Age <= 0 {
		err := fmt.Errorf("age must be positive, got: %d", u.Age)
		log.Printf("Validation error: %v", err)
		return err
	}

	if u.PhoneNumber == "" {
		err := errors.New("phone number is required")
		log.Printf("Validation error: %v", err)
		return err
	}

	if u.Email == "" {
		err := errors.New("email is required")
		log.Printf("Validation error: %v", err)
		return err
	}

	if !emailRegex.MatchString(u.Email) {
		err := fmt.Errorf("invalid email format: %s", u.Email)
		log.Printf("Validation error: %v", err)
		return err
	}

	log.Printf("User validation successful")
	return nil
} 