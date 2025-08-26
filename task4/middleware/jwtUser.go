package middleware

import (
	"net/http"
	"task4/config"
	"task4/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ContextUserIDKey  = "userID"
	ContextPostKey    = "post"
	ContextCommentKey = "comment"
)

func JWTUserMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 从请求头获取Authorization
		headers := ctx.Request.Header
		logger.Log.Debugf("请求头: %+v", headers)
		tokenString := ctx.GetHeader("Token")
		if tokenString == "" {
			logger.Log.Warn("缺少认证头信息")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10001,
				"message": "缺少认证令牌",
			})
			return
		}

		// 3. 验证并解析JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrHashUnavailable
			}
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			logger.Log.Warnf("令牌验证失败: %v", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10003,
				"message": "无效的认证令牌",
			})
			return
		}

		// 4. 提取用户ID
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, ok := claims["user_id"].(float64); ok {
				logger.Log.Debugf("认证成功，用户ID: %v", uint(userID))
				ctx.Set(ContextUserIDKey, uint(userID))
				ctx.Next()
				return
			}
		}

		// 5. 令牌有效但无法解析claims
		logger.Log.Warn("令牌缺少用户声明")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    10004,
			"message": "令牌格式错误",
		})

	}
}
