package model

import "time"

type ServiceRetry struct {
	ID        int64      `json:"id" form:"id" gorm:"column:id"`                                      // 自增ID
	BizID     int64      `json:"biz_id" form:"biz_id" gorm:"column:biz_id"`                          // 重试任务ID
	Type      int64      `json:"type" form:"type" gorm:"column:type"`                                // 任务状态
	Data      string     `json:"data" form:"data" gorm:"column:data"`                                // 序列化数据json
	Status    int64      `json:"status" form:"status" gorm:"column:status"`                          // 任务状态: 0未完成 1完成 2超过次数放弃
	RetryNum  int64      `json:"retry_num" form:"retry_num" gorm:"column:retry_num"`                 // 已重试次数
	TraceID   string     `json:"trace_id" form:"trace_id" gorm:"column:trace_id"`                    // 此次任务的trace
	CreatedAt *time.Time `json:"created_at" form:"created_at" gorm:"column:created_at"`              // 创建时间, 防止nil无法解析，需改为*time.Time
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at" gorm:"column:updated_at;default:null"` // 更新时间, 时间添加默认空值，防止time.Time空值出错
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at" gorm:"column:deleted_at;default:null"` // 删除时间
}
