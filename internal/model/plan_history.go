package model

import (
	"time"
)

// PlanHistory 计划历史记录模型
type PlanHistory struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PlanID     uint      `gorm:"not null;index" json:"plan_id"`
	FieldName  string    `gorm:"size:50;not null" json:"field_name"`
	OldValue   string    `gorm:"type:text" json:"old_value"`
	NewValue   string    `gorm:"type:text" json:"new_value"`
	ChangeType string    `gorm:"size:20;not null" json:"change_type" binding:"required,oneof=Status Progress Info"`
	ChangedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"changed_at"`
}

// TableName 指定表名
func (PlanHistory) TableName() string {
	return "plan_histories"
}

// ChangeType 变更类型常量
const (
	ChangeTypeStatus   = "Status"
	ChangeTypeProgress = "Progress"
	ChangeTypeInfo     = "Info"
)
