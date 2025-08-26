package middleware

import (
	"net/http"
	"strconv"

	"task4/database"
	"task4/logger"
	"task4/models"

	"github.com/gin-gonic/gin"
)

// PostOwnershipMiddleware 验证当前用户是否是文章作者
func PostOwnershipMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 获取当前用户ID
		userID, exists := ctx.Get("userID")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10005,
				"message": "用户未认证",
			})
			return
		}

		// 2. 获取文章ID
		postID := ctx.Param("id")

		// 3. 查询文章
		var post models.Post
		if err := database.DB.First(&post, postID).Error; err != nil {
			logger.Log.Warnf("文章查询失败: %v", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}

		// 4. 验证所有者
		if post.UserID != userID.(uint) {
			logger.Log.Warnf("用户 %d 尝试操作未授权文章 %d", userID, post.ID)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "无权操作此文章",
			})
			return
		}

		ctx.Set(ContextPostKey, &post)
		ctx.Next()
	}
}

// CommentOwnershipMiddleware 验证当前用户是否是文章作者
func CommentOwnershipMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 从上下文获取用户ID
		userID, exists := ctx.Get("userID")
		if !exists {
			logger.Log.Warn("用户未认证")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    10001,
				"message": "用户未认证",
			})
			return
		}

		// 2. 从URL获取评论ID
		commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			logger.Log.Warnf("无效的评论ID: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    20001,
				"message": "无效的评论ID",
			})
			return
		}

		// 3. 查询评论
		var comment models.Comment
		if err := database.DB.First(&comment, commentID).Error; err != nil {
			logger.Log.Warnf("评论查询失败: %v", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "评论不存在",
			})
			return
		}

		// 4. 验证所有者
		if comment.UserID != userID.(uint) {
			logger.Log.Warnf("用户 %d 尝试操作未授权评论 %d", userID, commentID)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "无权操作此评论",
			})
			return
		}

		// 将评论对象存储到上下文，供后续处理使用
		ctx.Set(ContextCommentKey, &comment)
		ctx.Next()
	}
}
