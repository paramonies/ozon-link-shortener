package utils

import (
	"crypto/sha256"
	"math/big"
	"regexp"
)

// All characters
const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = len(alphabet)
)

func sha256Hash(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

func Convert(id int, link string) string {
	idNot62 := big.NewInt(int64(id)).Text(base)
	urlHashBytes := sha256Hash(link + string(idNot62))
	result := new(big.Int).SetBytes(urlHashBytes).Text(base)
	return result[:10]
}

func IsShortValid(id string) bool {
	pattern := `^[0-9a-zA-z]{10}$`
	isValid, err := regexp.Match(pattern, []byte(id))
	if err != nil {
		return false
	}
	return isValid
}
