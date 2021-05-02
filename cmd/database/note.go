package database

import "time"

type Note struct {
	ID         uint       `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt  time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsListNote bool       `json:"is_list_note"`
	IsPublic   bool       `json:"is_public"`
	OwnerID    uint       `json:"owner_id" sql:"type:int REFERENCES users(id)"`
	Owner      User       `json:"owner"`
	FolderID   uint       `json:"folder_id"`
	NoteBodies []NoteBody `json:"note_bodies"`
	Name       string     `json:"name" gorm:"size:255;not null;""`
}
