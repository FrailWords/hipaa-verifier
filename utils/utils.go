package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1HashOfValue(value string) string {
	h := sha1.New()
	h.Write([]byte(value))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}
