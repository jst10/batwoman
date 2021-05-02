package database

type UserNote struct {
	ID        uint    `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	UserID   uint   `json:"user_id" sql:"type:int REFERENCES users(id)"`
	User     User   `json:"user"`
	NoteID   uint   `json:"note_id" sql:"type:int REFERENCES notes(id)"`
	Note     User   `json:"note"`
}
