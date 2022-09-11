package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

func generateRandomSalt() []byte {
	var salt = make([]byte, 16)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}

func hashValueWithSalt(value string, salt []byte) string {
	var valueBytes = []byte(value)
	var sha512Hasher = sha512.New()
	valueBytes = append(valueBytes, salt...)
	sha512Hasher.Write(valueBytes)
	var hashedValueBytes = sha512Hasher.Sum(nil)
	var hashedValueHex = hex.EncodeToString(hashedValueBytes)
	return hashedValueHex
}

func Sha512HashOfValue(value string) string {
	var salt = generateRandomSalt()
	return hashValueWithSalt(value, salt)
}
