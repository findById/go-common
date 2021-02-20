package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HashMD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HashSHA1(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HashSHA256(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HashSHA512(b []byte) string {
	h := sha512.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HashMD5String(b string) string {
	return HashMD5([]byte(b))
}

func HashSHA1String(b string) string {
	return HashSHA1([]byte(b))
}

func HashSHA256String(b string) string {
	return HashSHA256([]byte(b))
}

func HashSHA512String(b string) string {
	return HashSHA512([]byte(b))
}

func HmacSHA1String(msg, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSHA256String(msg, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSHA512String(msg, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}
