package passwordGenerator

import (
	"fmt"
	"math/rand/v2"
	"unicode"
)

func passGen(length int) string {
	var password string
	characters := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ !#$%&()*+-/:;<=>?@[]^_{|}~"
	if length < 3 {
		fmt.Println("Not enough characters")
		return password
	}
	for len(password) < length {
		rndChar := string(characters[rand.IntN(89)])
		if unicode.IsNumber(rune(rndChar[0])) {
			continue
		}

		password += rndChar
	}

	return password
}
