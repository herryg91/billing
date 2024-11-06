package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

const randomStringLetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	randomStringletterIdxBits = 6
	randomStringletterIdxMask = 1<<randomStringletterIdxBits - 1
	randomStringletterIdxMax  = 63 / randomStringletterIdxBits
)

func RandomString(n int) string {
	// from : https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), randomStringletterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), randomStringletterIdxMax
		}
		if idx := int(cache & randomStringletterIdxMask); idx < len(randomStringLetterBytes) {
			b[i] = randomStringLetterBytes[idx]
			i--
		}
		cache >>= randomStringletterIdxBits
		remain--
	}

	return string(b)
}

const letterNumberBytes = "0123456789"

func RandomStringIntOnly(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterNumberBytes[rand.Int63()%int64(len(letterNumberBytes))]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func GetOrdinalNumber(n int) string {
	if n >= 11 && n <= 13 {
		return fmt.Sprintf("%dth", n)
	}

	switch n % 10 {
	case 1:
		return fmt.Sprintf("%dst", n)
	case 2:
		return fmt.Sprintf("%dnd", n)
	case 3:
		return fmt.Sprintf("%drd", n)
	default:
		return fmt.Sprintf("%dth", n)
	}
}

func StringCoalesce(in string, fallback string) string {
	if in == "" {
		return fallback
	}
	return in
}
