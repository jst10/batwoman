package database

import "time"

type Session struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uint      `json:"user_id" sql:"type:int REFERENCES users(id)"`
	User      User      `json:"owner"`
}
