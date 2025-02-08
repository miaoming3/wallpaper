package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/miaoming3/wallpaper/app/global"
	"io"
)

func Encrypt(txt []byte) (string, error) {

	block, err := aes.NewCipher([]byte(global.SysConfig.AesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, txt, nil)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt() {

}
