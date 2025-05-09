package bcrypt

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password using bcrypt with a cost of 10
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a password with a hashed password and returns an error if they do not match
// It returns nil if they match
func CheckPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
