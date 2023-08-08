package Aes

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/yibana/hackWallet/configs"
	"log"
	"os"
	"strings"
	"testing"
)

func init() {
	println(configs.PROXY)
}

func TestEncrypt(t *testing.T) {
	key, _ := hex.DecodeString(configs.AESKEY)
	//input privateKey from clipboard
	privateKey, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalf("无法读取剪贴板内容：%v", err)
	}
	encryptedText, err := Encrypt(key, strings.Trim(privateKey, "\n"))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Encrypted privateKey:", encryptedText)
}

func TestInit(t *testing.T) {
	key := make([]byte, 16)
	//key := []byte("my-secret-key-123") // 用于加密和解密的密钥
	_, err := rand.Read(key)
	fmt.Printf("key:%x\n", key)
	// 加密文本
	encryptedText, err := Encrypt(key, os.Getenv("privateKey"))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Encrypted privateKey:", encryptedText)

	encryptedText, err = Encrypt(key, os.Getenv("signingKey"))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Encrypted signingKey:", encryptedText)

	// 解密文本
	decryptedText, err := Decrypt(key, encryptedText)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted text:", decryptedText)
}
func TestDecrypt(t *testing.T) {
	key, _ := hexutil.Decode("")

	// 解密文本
	decryptedText, err := Decrypt(key, os.Getenv("privateKey"))
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted privateKey:", decryptedText)

	decryptedText, err = Decrypt(key, os.Getenv("signingKey"))
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted signingKey:", decryptedText)
}
