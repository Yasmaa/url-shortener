package alias

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase62(number int) string {
	if number == 0 {
		return string(alphabet[0])
	}
	encoded := ""
	base := len(alphabet)
	for number > 0 {
		remainder := number % base
		encoded = string(alphabet[remainder]) + encoded
		number = number / base
	}
	return encoded
}

func GenerateBase62ID(length int) string {
	var result string
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		result += string(alphabet[randomIndex])
	}
	return result + EncodeBase62(int(time.Now().UnixNano()))
}
