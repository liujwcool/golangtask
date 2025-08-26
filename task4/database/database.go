package database

import (
	"fmt"
	"task4/config"
	"task4/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)
	fmt.Println("dsn = ", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.Log.Errorf("数据库连接失败: %v", err)
		return fmt.Errorf("数据库连接失败: %w", err)
	}
	// 添加连接测试
	sqlDB, err := DB.DB()
	if err == nil {
		if pingErr := sqlDB.Ping(); pingErr != nil {
			logger.Log.Errorf("数据库连接测试失败: %v", err)
			return fmt.Errorf("数据库连接测试失败: %w", pingErr)
		}
	}

	logger.Log.Info("数据库连接成功")
	return nil
}
