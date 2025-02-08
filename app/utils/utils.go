package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

func ResponseJsonUnmarshal(data interface{}, responseData interface{}) error {
	imageJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(imageJson, responseData); err != nil {
		return err
	}
	return nil
}

func RemoteUrl(c *gin.Context, imgUrl string) string {
	scheme := c.Request.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		scheme = "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
	}

	return fmt.Sprintf("%v://%v/%v", scheme, c.Request.Host, imgUrl)
}
