package password

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(length int, digits, symbols bool) (string, error) {
	var passwordRunes []rune
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitRunes := []rune("0123456789")
	symbolRunes := []rune("!@#$%^&*()-_=+[]{};:,.<>?")

	if digits {
		passwordRunes = append(passwordRunes, digitRunes...)
	}
	if symbols {
		passwordRunes = append(passwordRunes, symbolRunes...)
	}
	passwordRunes = append(passwordRunes, letterRunes...)

	result := make([]rune, length)
	for i := range result {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordRunes))))
		if err != nil {
			return "", err
		}
		result[i] = passwordRunes[randIndex.Int64()]
	}
	return string(result), nil
}
