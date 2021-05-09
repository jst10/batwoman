package database

import (
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"time"
)

type Folder struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	OwnerID   uint      `json:"owner_id" gorm:"not null" sql:"type:int REFERENCES users(id)"`
	Owner     User      `json:"owner"  gorm:"constraint:OnDelete:CASCADE"`
	// TODO if there will be time for this
	//ParentFolderID *uint     `json:"parent_folder_id"`
	//ParentFolder   *Folder   `json:"parent_folder"`
	Notes []Note `json:"notes"`
	Name  string `json:"name" gorm:"size:255;not null;"`
}

func (item *Folder) Create() (*Folder, *custom_errors.CustomError) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	result := db.Debug().Create(&item)
	if result.Error != nil {
		return &Folder{}, custom_errors.GetDbError(result.Error, getType(item)+"->Create")
	}
	return item, nil
}

func (item *Folder) Update(uid uint) (*Folder, *custom_errors.CustomError) {
	result := db.Debug().Model(&Folder{}).Where("id = ?", uid).Take(&Folder{}).UpdateColumns(
		map[string]interface{}{
			"name":       item.Name,
			"updated_at": time.Now(),
		},
	)
	if result.Error != nil {
		return &Folder{}, custom_errors.GetDbError(result.Error, getType(item)+"->Update")
	}
	return item.GetByID(uid)
}


func (item *Folder) List(queryOptions *FolderQueryOptions) (*PageOfFolders, *custom_errors.CustomError) {
	items := []Folder{}
	query := db.Debug().Model(&Folder{})

	conditions := ""
	conditionParams := []interface{}{}

	if len(queryOptions.OrderBy) > 0 {
		query = query.Order(queryOptions.OrderBy + " " + queryOptions.OrderDirection)
	}

	if len(queryOptions.OwnerId) > 0 {
		conditions = extendConditions(conditions, "owner_id=?")
		conditionParams = append(conditionParams, queryOptions.OwnerId)
	}
	if len(queryOptions.Name) > 0 {
		conditions = extendConditions(conditions, "name LIKE %?%")
		conditionParams = append(conditionParams, queryOptions.Name)
	}
	if len(conditions) > 0 {
		query = query.Where(conditions, conditionParams...)
	}

	limit := queryOptions.PageSize
	offset := (queryOptions.Page - 1) * queryOptions.PageSize
	var count int64
	query.Count(&count)
	query = query.Offset(offset).Limit(limit)
	result := query.Find(&items)
	if result.Error != nil {
		return nil, custom_errors.GetDbError(result.Error, getType(item)+"->All")
	}

	pageOfFolders := PageOfFolders{}
	pageOfFolders.Count = int(count)
	pageOfFolders.Page = queryOptions.Page
	pageOfFolders.PageSize = queryOptions.PageSize
	pageOfFolders.Items = items
	return &pageOfFolders, nil
}

func (item *Folder) GetByID(id uint) (*Folder, *custom_errors.CustomError) {
	result := db.Debug().Model(Folder{}).Where("id = ?", id).Take(&item)
	if result.Error != nil {
		return &Folder{}, custom_errors.GetDbError(result.Error, getType(item)+"->GetByID")
	}
	return item, nil
}

func (item *Folder) DeleteById(id uint) (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Folder{}).Where("id = ?", id).Delete(&Folder{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteById")
	}
	return result.RowsAffected, nil
}

func (item *Folder) DeleteAll() (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Folder{}).Where("1=1").Delete(&Folder{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteAll")
	}
	return db.RowsAffected, nil
}
