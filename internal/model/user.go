package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:text"` // SQLite推荐使用text类型
	Email     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
