package generate_password

import (
	"crypto/rand"
	"math/big"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	// TODO: ваш код
	if n > 0 {
		str := make([]byte, 0, n)
		for range make([]struct{}, n) {
			randNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err != nil {
				return "", nil
			}
			str = append(str, letters[randNum.Int64()])
		}
		return string(str), nil
	} else {
		return "", nil
	}

}
