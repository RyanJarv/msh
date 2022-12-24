package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetName(args []string) string {
	h := sha1.New()
	h.Write([]byte(strings.Join(args, "")))

	hex := fmt.Sprintf("%x", h.Sum(nil))

	cmd := strings.ReplaceAll(filepath.Base(args[0]), ".", "-")
	return fmt.Sprintf("%s-%s", cmd, hex)
}

func Sha256[T []byte | string](s T) T {
	sha := sha256.Sum256([]byte(s))
	return T(base64.StdEncoding.EncodeToString(sha[:]))
}

func Sha256Base64[T []byte | string](s T) T {
	return T(base64.StdEncoding.EncodeToString([]byte(s)))
}
