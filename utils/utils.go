package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
	userIDByte := []byte(userID)
	hash := md5.Sum(userIDByte)

	return strings.ToUpper(hex.EncodeToString(hash[:]))[:6]
}

func IsWalletNumberValid(userID, walletNumber string) bool {
	generatedWalletNumber := GenerateWalletNumber(userID)
	return generatedWalletNumber == walletNumber
}

func GenerateAppURL() string {
	settings := AppSettings.AppParams
	return fmt.Sprintf("http://%s:%s/api/%s", settings.ServerURL, settings.PortRun, settings.AppVersion)
}
