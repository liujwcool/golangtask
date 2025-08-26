package utils

import (
	"errors"
	"strings"
	"task4/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// logger.Log.Printf("输入密码: %q\n", password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// logger.Log.Printf("存储哈希: %q\n", string(bytes))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logger.Log.Printf("密码验证失败: %v\n", err)
		logger.Log.Printf("输入密码: %q\n", password)
		logger.Log.Printf("存储哈希: %q\n", hash)
		logger.Log.Printf("哈希长度: %d\n", len(hash))
		// hashpassword, err := HashPassword(password)
		// logger.Log.Printf("存储哈希: %q\n", hashpassword)
		// logger.Log.Printf("哈希长度: %d\n", len(hashpassword))

		// 检查常见错误模式
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			logger.Log.Println("错误: 密码不匹配")
		} else if strings.Contains(err.Error(), "hashedPassword") {
			logger.Log.Println("错误: 哈希格式无效")
		}
		return false
	}
	return true
}
