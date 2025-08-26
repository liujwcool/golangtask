package main

import (
	"fmt"
	"task4/config"
	"task4/database"
	"task4/logger"
	"task4/middleware"
	"task4/models"
	"task4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		// logger.Log.Fatalf("加载配置失败: %v", err)
		fmt.Printf("加载配置失败: %v", err)
		return
	}
	// logger.Log.Info("配置加载成功")
	fmt.Println("配置加载成功")

	logger.InitLogger(&cfg)

	database.InitDB(&cfg)
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		logger.Log.Fatalf("数据库迁移失败: %v", err)
	}
	logger.Log.Info("数据库迁移完成")

	router := gin.New()

	router.Use(middleware.RecoveryMiddleware())

	// router.Use(middleware.JWTUserMiddleware(&cfg))

	routes.SetupRoutes(router, &cfg)

	logger.Log.Infof("服务器启动，监听端口 %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Log.Fatalf("服务启动失败: %v", err)
	}

}
