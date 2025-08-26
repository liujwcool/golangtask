package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关系定义
	UserID   uint      `gorm:"index;not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"author,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}
