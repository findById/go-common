package aes_util

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAESCBCEncrypt(t *testing.T) {
	key := "0000000000000000"
	data := "Hello, world!"

	fmt.Println("plainText:", data)

	b, err := AESCBCEncrypt([]byte(data), []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println("encryptText:", hex.EncodeToString(b))

	b, err = AESCBCDecrypt(b, []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println("decryptText:", string(b))
}
