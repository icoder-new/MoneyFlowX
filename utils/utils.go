package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"reflect"
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

func CalculateHash(item interface{}) string {
	h := hmac.New(sha1.New, []byte(AppSettings.AppParams.SecretKey))

	var addBytes func(reflect.Value)
	addBytes = func(value reflect.Value) {
		switch value.Kind() {
		case reflect.Ptr, reflect.Interface:
			if value.IsNil() {
				h.Write([]byte("nil"))
			} else {
				addBytes(value.Elem())
			}
		case reflect.Slice, reflect.Array:
			for i := 0; i < value.Len(); i++ {
				addBytes(value.Index(i))
			}
		case reflect.Struct:
			for i := 0; i < value.NumField(); i++ {
				addBytes(value.Field(i))
			}
		default:
			itemBytes := make([]byte, 8)
			switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				binary.BigEndian.PutUint64(itemBytes, uint64(value.Int()))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				binary.BigEndian.PutUint64(itemBytes, uint64(value.Uint()))
			case reflect.Float32, reflect.Float64:
				bits := math.Float64bits(value.Float())
				binary.BigEndian.PutUint64(itemBytes, bits)
			case reflect.Bool:
				if value.Bool() {
					itemBytes[0] = 1
				} else {
					itemBytes[0] = 0
				}
			case reflect.String:
				strBytes := []byte(value.String())
				h.Write(strBytes)
				return
			}
			h.Write(itemBytes)
		}
	}

	itemValue := reflect.ValueOf(item)

	addBytes(itemValue)

	return fmt.Sprintf("%x", h.Sum(nil))
}
