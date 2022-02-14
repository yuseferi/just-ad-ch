package helpers

import (
	"crypto/md5"
	"fmt"
)

// MD5 hashes using md5 algorithm
func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}
