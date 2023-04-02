package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPassword(password, hash string) bool {
	fmt.Println(password, "MMMMcheckpasswordMMMM")
	fmt.Println(hash, "checkhash")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err, "<<<<er")
	return err == nil
}
