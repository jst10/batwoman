package database

type User struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
