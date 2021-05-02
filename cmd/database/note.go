package database

type Note struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
	IsListNote bool       `json:"is_list_note"`
	OwnerID    uint       `json:"owner_id"`
	Owner      User       `json:"owner"`
	FolderID   uint       `json:"folder_id"`
	NoteBodies []NoteBody `json:"note_bodies"`
	Name       string     `json:"name"`
}
