package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"task4/database"
	"task4/logger"
	"task4/middleware"
	"task4/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}

// CreatePost 创建文章
func (c *PostController) CreatePost(ctx *gin.Context) {
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

	var createReq struct {
		Title   string `json:"title" binding:"required,min=2,max=100"`
		Content string `json:"content" binding:"required,min=10"`
	}

	if err := ctx.ShouldBindJSON(&createReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 40001,
			"msg":  "请求参数无效",
		})
		return
	}

	newPost := models.Post{
		Title:   createReq.Title,
		Content: createReq.Content,
		UserID:  userID,
	}

	if err := database.DB.Create(&newPost); err.Error != nil {
		logger.Log.Errorf("创建文章失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50001,
			"msg":  "创建文章失败",
		})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "文章创建成功",
		"data": gin.H{
			"id":         newPost.ID,
			"title":      newPost.Title,
			"created_at": newPost.CreatedAt,
		},
	})
}

// GetPosts 获取文章列表
func (c *PostController) GetPosts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	var posts []models.Post
	result := database.DB.Order("created_at DESC").Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		logger.Log.Errorf("获取文章列表失败: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50001,
			"msg":  "获取文章列表失败",
		})
		return
	}

	var count int64
	database.DB.Model(&models.Post{}).Count(&count)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{
			"posts": posts,
			"pagination": gin.H{
				"current_page": page,
				"per_page":     limit,
				"total_items":  count,
				"total_pages":  (count + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// GetPost 获取单篇文章
func (c *PostController) GetPost(ctx *gin.Context) {
	postID := ctx.Param("id")

	var post models.Post
	result := database.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).First(&post, postID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code": 40401,
				"msg":  "文章不存在",
			})
			return
		}

		logger.Log.Errorf("获取文章详情失败: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50002,
			"msg":  "获取文章详情失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": post,
	})
}

// UpdatePost 更新文章
func (c *PostController) UpdatePost(ctx *gin.Context) {
	postObj, exists := ctx.Get(middleware.ContextPostKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50003,
			"msg":  "内部错误：未找到文章",
		})
		return
	}

	post, ok := postObj.(*models.Post)

	if !ok {
		logger.Log.Errorf("文章对象类型无效: %T", postObj)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50004,
			"msg":  "内部错误：无效的文章对象",
		})
		return
	}

	var updateReq struct {
		Title   string `json:"title" binding:"required,min=2,max=100"`
		Content string `json:"content" binding:"required,min=10"`
	}

	if err := ctx.ShouldBindJSON(&updateReq); err.Error() == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  40001,
			"msg":   "请求参数无效",
			"error": err.Error(),
		})
		return
	}

	post.Title = updateReq.Title
	post.Content = updateReq.Content

	result := database.DB.Save(&post)
	if result.Error != nil {
		logger.Log.Errorf("更新文章失败: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50005,
			"msg":  "更新文章失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "文章更新成功",
		"data": gin.H{
			"id":         post.ID,
			"title":      post.Title,
			"updated_at": post.UpdatedAt,
		},
	})

}

// DeletePost 删除文章
func (c *PostController) DeletePost(ctx *gin.Context) {

	postObj, exists := ctx.Get(middleware.ContextPostKey)

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50006,
			"msg":  "内部错误：未找到文章",
		})
		return
	}

	// 类型断言
	post, ok := postObj.(*models.Post)
	if !ok {
		logger.Log.Errorf("文章对象类型无效: %T", postObj)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50007,
			"msg":  "内部错误：无效的文章对象",
		})
		return
	}

	result := database.DB.Delete(&post)
	if result.Error != nil {
		logger.Log.Errorf("删除文章失败: %v", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50008,
			"msg":  "删除文章失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "文章删除成功",
	})
}
