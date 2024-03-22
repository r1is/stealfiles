package common

import (
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/emmansun/gmsm/padding"
	"github.com/emmansun/gmsm/sm3"
	"github.com/emmansun/gmsm/sm4"
)

func Sm4_d(_key string, _ciphertext string) string {
	strKey := GetSm3KeyToSm4(_key)
	key, _ := hex.DecodeString(strKey)
	ciphertext, _ := hex.DecodeString(_ciphertext)

	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < sm4.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:sm4.BlockSize]
	ciphertext = ciphertext[sm4.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad plaintext
	pkcs7 := padding.NewPKCS7Padding(sm4.BlockSize)
	ciphertext, err = pkcs7.Unpad(ciphertext)
	if err != nil {
		fmt.Println("runing error")
		panic(err)
	}

	return string(ciphertext)
}

// 直接使用sm3.Sum方法

func GetSm3KeyToSm4(key string) string {
	sum := sm3.Sum([]byte(key))
	a := hex.EncodeToString(sum[:])
	return a[32:]
}
