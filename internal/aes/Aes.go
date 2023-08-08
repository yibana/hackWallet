package Aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

func Encrypt(key []byte, text string) (string, error) {
	// 创建一个 AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 生成一个随机的初始化向量（IV）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用 CBC 模式加密文本
	mode := cipher.NewCBCEncrypter(block, iv)
	padded := PKCS5Padding([]byte(text), aes.BlockSize)
	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	// 将 IV 和密文进行 Base64 编码
	encodedIV := base64.StdEncoding.EncodeToString(iv)
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	// 返回 IV 和密文的 Base64 编码
	return encodedIV + ":" + encodedCiphertext, nil
}

func Decrypt(key []byte, ciphertext string) (string, error) {
	// 分离出 IV 和密文
	parts := strings.Split(ciphertext, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid ciphertext format")
	}
	encodedIV, encodedCiphertext := parts[0], parts[1]

	// 对 IV 和密文进行 Base64 解码
	iv, err := base64.StdEncoding.DecodeString(encodedIV)
	if err != nil {
		return "", err
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", err
	}

	// 创建一个 AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 使用 CBC 模式解密文本
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(decodedCiphertext))
	mode.CryptBlocks(plaintext, decodedCiphertext)

	// 去除填充（padding）
	unpadded, err := PKCS5Unpadding(plaintext)
	if err != nil {
		return "", err
	}

	// 返回解密后的明文
	return string(unpadded), nil
}

func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS5Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}
