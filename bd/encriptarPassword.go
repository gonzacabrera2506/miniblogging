package bd

import "golang.org/x/crypto/bcrypt"

//EncriptarPasswprd permite encriptar la contraseña recibida
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
