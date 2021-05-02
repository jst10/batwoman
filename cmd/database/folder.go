package database

import "time"

type Folder struct {
	ID             uint      `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	OwnerID        uint      `json:"owner_id" sql:"type:int REFERENCES users(id)"`
	Owner          User      `json:"owner"`
	ParentFolderID *uint     `json:"parent_folder_id"`
	ParentFolder   *Folder   `json:"parent_folder"`
	Notes          []Note    `json:"notes"`
	Name           string    `json:"name" gorm:"size:255;not null;"`
}
