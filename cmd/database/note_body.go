package database

import "time"

type NoteBody struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	NoteID    uint      `json:"note_id" gorm:"not null;" sql:"type:int REFERENCES notes(id)"`
	Text      string    `json:"text" gorm:"size:2048;not null;"`
}
