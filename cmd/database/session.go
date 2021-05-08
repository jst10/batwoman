package database

import (
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"time"
)

type Session struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uint      `json:"user_id" sql:"type:int REFERENCES users(id)"`
	User      User      `json:"owner"`
}

func (item *Session) Create() (*Session, *custom_errors.CustomError) {
	result := db.Debug().Create(&item)
	if result.Error != nil {
		return &Session{}, custom_errors.GetDbError(result.Error, getType(item)+"->Create")
	}
	return item, nil
}

func (item *Session) DeleteById(id uint) (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Session{}).Where("id = ?", id).Take(&Session{}).Delete(&Session{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteById")
	}
	return result.RowsAffected, nil
}

func (item *Session) DeleteByUserId(userId uint) (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Session{}).Where("userid = ?", userId).Delete(&Session{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteByUserId")
	}
	return result.RowsAffected, nil
}

func (item *Session) GetByID(id uint) (*Session, *custom_errors.CustomError) {
	result := db.Debug().Model(Session{}).Where("id = ?", id).Take(&item)
	if result.Error != nil {
		return &Session{}, custom_errors.GetDbError(result.Error, getType(item)+"->GetByID")
	}
	return item, nil
}
