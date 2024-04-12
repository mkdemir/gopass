package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GeneratePassword(length int, digits, symbols, upper bool) (string, error) {
	var passwordRunes []rune
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	digitRunes := []rune("0123456789")
	symbolRunes := []rune("!@#$%^&*()-_=+[]{};:,.<>?")

	if upper {
		letterRunes = append(letterRunes, []rune(strings.ToUpper(string(letterRunes)))...)
	}

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
