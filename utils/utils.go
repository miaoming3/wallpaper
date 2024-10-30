package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// Md5 返回加密后字符串
func Md5(data string) string {
	byteData := []byte(data)
	hash := md5.Sum(byteData)
	return hex.EncodeToString(hash[:])
}

// GenerateRandomStringMath 返回指定长度随机字符串
func GenerateRandomStringMath(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

// GenerateRandomBytes 返回指定长度随机字节
func GenerateRandomBytes(n int) []byte {
	result := make([]byte, n)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = byte(seededRand.Intn(256)) // 生成0-255之间的随机字节
	}
	return result
}
