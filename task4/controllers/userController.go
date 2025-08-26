package controllers

import (
	"net/http"
	"task4/config"
	"task4/database"
	"task4/models"
	"task4/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Cfg *config.Config
}

func NewUserController(cfg *config.Config) *UserController {
	return &UserController{Cfg: cfg}
}

// 注册
func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 40001,
			"msg":  "请求参数错误",
		})
		return
	}

	// 查询username是否存在
	var exists int64
	database.DB.Model(&models.User{}).Where("username = ?", user.Username).Count(&exists)

	if exists > 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"code": 40002,
			"msg":  "用户名已被占用",
		})
		return
	}
	//

	if result := database.DB.Create(&user); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50001,
			"msg":  "注册失败",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"username": user.Username,
			"email":    user.Email,
			"user_id":  user.ID,
		},
	})
}

// 登录
func (c *UserController) Login(ctx *gin.Context) {

	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 40003,
			"msg":  "用户名或密码不能为空",
		})
		return
	}

	var user models.User
	if result := database.DB.Model(&models.User{}).
		Where("username = ?", credentials.Username).
		First(&user); result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 40004,
			"msg":  "用户名或密码错误",
		})
		return
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 40005,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, c.Cfg.JWTSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 50002,
			"msg":  "登录失败，请重试",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"token":   token,
			"user_id": user.ID,
			"expire":  time.Now().Add(time.Hour * 24).Format(time.RFC3339),
		},
	})

}
