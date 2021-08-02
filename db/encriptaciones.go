package db

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword retorna el string encriptado o error.
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	passBytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(passBytes), err
}
