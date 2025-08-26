package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关系定义
	UserID uint `gorm:"index;not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"author,omitempty"`

	PostID uint `gorm:"index;not null" json:"post_id"`
	Post   Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

// AfterDelete 钩子 - 删除评论后检查文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comments_status", "无评论").Error
	}
	return nil
}
