package model

import (
	"time"

	"gorm.io/gorm"
)

// Plan 计划模型
type Plan struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Priority    string         `gorm:"size:20;not null" json:"priority" binding:"required,oneof=High Medium Low"`
	Status      string         `gorm:"size:20;not null" json:"status" binding:"required,oneof=Todo InProgress Done Cancelled"`
	DueDate     *time.Time     `json:"due_date"`
	Progress    int            `gorm:"not null;default:0" json:"progress" binding:"min=0,max=100"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}

// TableName 指定表名
func (Plan) TableName() string {
	return "plans"
}

// BeforeCreate 创建前钩子
func (p *Plan) BeforeCreate(tx *gorm.DB) error {
	// 设置默认值
	if p.Status == "" {
		p.Status = "Todo"
	}
	if p.Progress == 0 {
		p.Progress = 0
	}
	return nil
}

// BeforeUpdate 更新前钩子
func (p *Plan) BeforeUpdate(tx *gorm.DB) error {
	// 如果进度达到 100%,自动更新状态为 Done
	if p.Progress == 100 && p.Status != "Done" {
		p.Status = "Done"
	}
	return nil
}
