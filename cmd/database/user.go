package database

import (
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Username  string `json:"username" gorm:"size:512;not null;unique"`
	Password  string `json:"password" gorm:"size:512;not null;"`
	Salt      string `json:"salt" gorm:"size:512;not null;"`
}

func (item *User) Create() (*User, *custom_errors.CustomError) {
	result := db.Debug().Create(&item)
	if result.Error != nil {
		return &User{}, custom_errors.GetDbError(result.Error, getType(item)+"->Create")
	}
	return item, nil
}

func (item *User) Update(uid uint) (*User, *custom_errors.CustomError) {
	result := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   item.Password,
			"username":   item.Username,
			"updated_at": time.Now(),
		},
	)
	if result.Error != nil {
		return &User{}, custom_errors.GetDbError(result.Error, getType(item)+"->Update")
	}
	return item.GetByID(uid)
}

func (item *User) All() (*[]User, *custom_errors.CustomError) {
	items := []User{}
	result := db.Debug().Model(&User{}).Limit(100).Find(&items)
	if result.Error != nil {
		return &[]User{}, custom_errors.GetDbError(result.Error, getType(item)+"->All")
	}
	return &items, nil
}

func (item *User) GetByID(id uint) (*User, *custom_errors.CustomError) {
	result := db.Debug().Model(User{}).Where("id = ?", id).Take(&item)
	if result.Error != nil {
		return &User{}, custom_errors.GetDbError(result.Error, getType(item)+"->GetByID")
	}
	return item, nil
}

func (item *User) GetByUsername(username string) (*User,  *custom_errors.CustomError) {
	result := db.Debug().Model(User{}).Where("username = ?", username).Take(&item)
	if result.Error != nil {
		return &User{}, custom_errors.GetDbError(result.Error, getType(item)+"->GetByUsername")
	}
	return item, nil
}

func (item *User) DeleteById(id uint) (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteById")
	}
	return result.RowsAffected, nil
}

func (item *User) DeleteAll() (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&User{}).Where("1=1").Delete(&User{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteAll")
	}
	return db.RowsAffected, nil
}
