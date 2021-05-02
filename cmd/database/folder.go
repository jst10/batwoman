package database

type Folder struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	OwnerID   uint   `json:"owner_id"`
	Owner     User   `json:"owner"`
	ParentFolderID *uint
	ParentFolder   *Folder
	Notes []Note `json:"notes"`
	Name string `json:"name"`
}
