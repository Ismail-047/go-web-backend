package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawPassword string) (hashedPassowrd string, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	return string(hashedBytes), err
}

func VerifyPassword(hashedPassowrd string, rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassowrd), []byte(rawPassword))
}