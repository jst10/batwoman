package database

type NoteBody struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	NoteID    uint   `json:"note_id"`
	Text   string `json:"text"`
}
