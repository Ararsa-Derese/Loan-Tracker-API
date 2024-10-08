package userutil

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"loan/domain"
	"regexp"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the password with the hashed password

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// ValidateEmail validates the email

func ValidateEmail(email string) bool {
	// Define a regular expression for validating an email address
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email matches the regex pattern
	return emailRegex.MatchString(email)
}

// ValidatePassword validates the password

func ValidatePassword(password string) bool {
	return len(password) >= 8
}
func GenerateOTP() string {
	secretLength := 8
	return gotp.RandomSecret(secretLength)
}

// A function that checks if a the logged in user can manipulate the target user.
func CanManipulateUser(claims *domain.JwtCustomClaims, user *domain.User, manip string) *domain.Response {
	// If the user is a regular user, they can only manipulate their own account.
	if claims.Role == "user" {
		if user.ID != claims.UserID {
			var message string
			if manip == "add" {
				message = "A User cannot add a new user"
			} else {
				message = "A User cannot " + manip + " another user"
			}

			return &domain.Response{
				Err: errors.New("unauthorized"),

				Message: message,
			}
		}

		return nil
	}

	// If the user is an admin, they can manipulate all users except root user and other admin users.
	if claims.Role == "admin" {
		if user.Role == "root" {
			return &domain.Response{
				Err: errors.New("forbidden"),

				Message: "Cannot " + manip + " root user",
			}
		}

		if user.Role == "admin" && claims.UserID != user.ID {
			return &domain.Response{
				Err: errors.New("unauthorized"),

				Message: "Admin cannot " + manip + " another admin user",
			}
		}
	}

	// If the user is a root user, they can manipulate all users.
	return nil
}

func GenerateDeviceFingerprint(ipAddress, userAgent string) string {
	// Combine IP address and User-Agent
	combined := fmt.Sprintf("%s|%s", ipAddress, userAgent)

	// Create a SHA-256 hash of the combined string
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashBytes := hash.Sum(nil)

	// Convert hash to a hexadecimal string
	fingerprint := hex.EncodeToString(hashBytes)

	return fingerprint
}
