package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateString(length int) string {
	return stringWithCharset(length, charset)
}

func GenerateWalletNumber(userID string) string {
	randomNumber := rand.Intn(9999)
	hash := md5.Sum([]byte(userID))
	hashStr := fmt.Sprintf("%x", hash)
	walletNumber := hashStr[:8] + fmt.Sprintf("%04d", randomNumber)

	return walletNumber
}
