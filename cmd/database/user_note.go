package database

type UserNote struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
	NoteID   uint   `json:"note_id"`
	Note     User   `json:"note"`
}
