package controllers

import (
	"net/http"
	"strconv"

	"task4/database"
	"task4/logger"
	"task4/middleware"
	"task4/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// CreateComment 创建评论
func (c *CommentController) CreateComment(ctx *gin.Context) {
	currentUserID, exists := ctx.Get(middleware.ContextUserIDKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 10001,
			"msg":  "用户未认证",
		})
		return
	}
	// 转换userid
	userID, ok := currentUserID.(uint)
	if !ok {
		logger.Log.Errorf("用户ID类型无效: %T", currentUserID)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50001,
			"msg":  "内部错误",
		})
		return
	}

	postID := ctx.Param("id")

	// 验证文章ID格式
	postIDUint, err := strconv.ParseUint(postID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 40001,
			"msg":  "无效的文章ID",
		})
		return
	}

	var commentReq struct {
		Content string `json:"content" binding:"required,min=5,max=500"`
	}

	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  40002,
			"msg":   "评论内容无效",
			"error": err.Error(),
		})
		return
	}

	var post models.Post
	if err := database.DB.First(&post, postIDUint).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code": 40401,
				"msg":  "文章不存在",
			})
			return
		}

		logger.Log.Errorf("查询文章失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50001,
			"msg":  "创建评论失败",
		})
		return
	}

	newComment := models.Comment{
		Content: commentReq.Content,
		UserID:  userID,
		PostID:  uint(postIDUint),
	}

	if err := database.DB.Create(&newComment).Error; err != nil {
		logger.Log.Errorf("创建评论失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50002,
			"msg":  "创建评论失败",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "评论创建成功",
		"data": gin.H{
			"id":         newComment.ID,
			"content":    newComment.Content,
			"created_at": newComment.CreatedAt,
		},
	})
}

// GetComments 获取文章评论
func (c *CommentController) GetComments(ctx *gin.Context) {
	postID := ctx.Param("id")

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	// 校验分页参数
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	var comments []models.Comment
	result := database.DB.
		Where("post_id = ?", postID).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username") // 只加载必要的用户字段
		}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments)

	if result.Error != nil {
		logger.Log.Errorf("获取评论失败: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50003,
			"msg":  "获取评论失败",
		})
		return
	}

	var totalCount int64
	database.DB.Model(&models.Comment{}).
		Where("post_id = ?", postID).
		Count(&totalCount)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{
			"comments": comments,
			"pagination": gin.H{
				"current_page": page,
				"per_page":     limit,
				"total_items":  totalCount,
				"total_pages":  (totalCount + int64(limit) - 1) / int64(limit),
			},
		},
	})
}
