package database

import "time"

type UserNote struct {
	ID        uint   `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	UserID   uint   `json:"user_id" sql:"type:int REFERENCES users(id)"`
	User     User   `json:"user"`
	NoteID   uint   `json:"note_id" sql:"type:int REFERENCES notes(id)"`
	Note     User   `json:"note"`
}
