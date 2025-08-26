package models

import (
	"task4/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Email     string         `gorm:"size:100;not null" json:"email"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关系定义
	Posts    []Post    `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// BeforeCreate 钩子 - 创建用户前加密密码
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 使用 utils 包中的密码加密函数
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// CheckPassword 验证密码是否正确
func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
