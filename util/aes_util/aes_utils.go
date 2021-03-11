package aes_util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// 使用PKCS5进行填充
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AESCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	blockSize := block.BlockSize()
	rawData = PKCS5Padding(rawData, blockSize)

	// 初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	// block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	// block大小和初始向量大小一定要一致
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AESCBCDecrypt(encryptText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(encryptText) < blockSize {
		return nil, errors.New("encryptText too short")
	}
	iv := encryptText[:blockSize]

	encryptText = encryptText[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptText)%blockSize != 0 {
		return nil, errors.New("encryptText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptText, encryptText)
	// 解填充
	encryptText = PKCS5UnPadding(encryptText)
	return encryptText, nil
}
