package util

import (
	"crypto/aes"
	"encoding/hex"
)

func AESEncrypt(key, value string) (string, error) {
	b, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	var dist []byte
	b.Encrypt(dist, []byte(value))
	return hex.EncodeToString(dist), nil
}

func AESDecrypt(key, value string) (string, error) {
	b, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	var dist []byte
	b.Decrypt(dist, []byte(value))
	return hex.EncodeToString(dist), nil
}
