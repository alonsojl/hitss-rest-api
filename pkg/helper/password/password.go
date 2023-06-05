package password

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	coste := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), coste)
	return string(hash), err
}

func Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
