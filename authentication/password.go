package authentication

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/mia0x75/pages/database"
)

// LoginIsCorrect TODO
func LoginIsCorrect(name string, password string) bool {
	hashedPassword, err := database.RetrieveHashedPasswordForUser([]byte(name))
	if len(hashedPassword) == 0 || err != nil { // len(hashedPassword) == 0 probably not needed.
		// User name likely doesn't exist
		return false
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

// EncryptPassword TODO
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
