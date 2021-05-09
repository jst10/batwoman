package database

import (
	"fmt"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"time"
)

type Note struct {
	ID         uint       `json:"id" gorm:"primaryKey;auto_increment"`
	CreatedAt  time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"not null"`
	IsListNote bool       `json:"is_list_note"`
	IsPublic   bool       `json:"is_public"`
	OwnerID    uint       `json:"owner_id" sql:"type:int REFERENCES users(id)"`
	Owner      User       `json:"owner" gorm:"constraint:OnDelete:CASCADE"`
	FolderID   uint       `json:"folder_id" sql:"type:int REFERENCES folders(id)"`
	NoteBodies []NoteBody `json:"note_bodies" gorm:"constraint:OnDelete:CASCADE"`
	Name       string     `json:"name" gorm:"size:255;not null;"`
}

func (item *Note) Create() (*Note, *custom_errors.CustomError) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	result := db.Debug().Create(&item)
	if result.Error != nil {
		return &Note{}, custom_errors.GetDbError(result.Error, getType(item)+"->Create")
	}
	return item, nil
}

func (item *Note) Update(uid uint) (*Note, *custom_errors.CustomError) {
	result := db.Debug().Model(&Note{}).Where("id = ?", uid).Take(&Note{}).UpdateColumns(
		map[string]interface{}{
			"name":       item.Name,
			"is_public":  item.IsPublic,
			"updated_at": time.Now(),
		},
	)

	if result.Error != nil {
		return &Note{}, custom_errors.GetDbError(result.Error, getType(item)+"->Update")
	}

	err := db.Debug().Model(item).Association("NoteBodies").Replace(item.NoteBodies)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return &Note{}, custom_errors.GetDbError(err, getType(item)+"->Update")
	}

	return item.GetByID(uid)
}

func (item *Note) List(queryOptions *NoteQueryOptions) (*PageOfNotes, *custom_errors.CustomError) {
	items := []Note{}
	query := db.Debug().Model(&Note{})

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
	if len(queryOptions.FolderId) > 0 {
		conditions = extendConditions(conditions, "folder_id=?")
		conditionParams = append(conditionParams, queryOptions.FolderId)
	}
	if len(queryOptions.SharedOption) > 0 {
		conditions = extendConditions(conditions, "is_public=?")
		conditionParams = append(conditionParams, queryOptions.SharedOption)
	} else {
		if queryOptions.OrPublic {
			if len(conditions) > 0 {
				conditions = "(" + conditions + ") or is_public=true"
			}
		}
	}

	if len(conditions) > 0 {
		query = query.Where(conditions, conditionParams...)
	}

	limit := queryOptions.PageSize
	offset := (queryOptions.Page - 1) * queryOptions.PageSize
	var count int64
	query.Count(&count)
	query = query.Offset(offset).Limit(limit)
	query=query.Preload("NoteBodies")
	result := query.Find(&items)
	if result.Error != nil {
		return nil, custom_errors.GetDbError(result.Error, getType(item)+"->All")
	}

	pageOfNotes := PageOfNotes{}
	pageOfNotes.Count = int(count)
	pageOfNotes.Page = queryOptions.Page
	pageOfNotes.PageSize = queryOptions.PageSize
	pageOfNotes.Items = items
	return &pageOfNotes, nil
}

func (item *Note) GetByID(id uint) (*Note, *custom_errors.CustomError) {
	result := db.Debug().Model(Note{}).Where("id = ?", id).Take(&item)
	if result.Error != nil {
		return &Note{}, custom_errors.GetDbError(result.Error, getType(item)+"->GetByID")
	}
	return item, nil
}

func (item *Note) DeleteById(id uint) (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Note{}).Where("id = ?", id).Delete(&Note{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteById")
	}
	return result.RowsAffected, nil
}

func (item *Note) DeleteAll() (int64, *custom_errors.CustomError) {
	result := db.Debug().Model(&Note{}).Where("1=1").Delete(&Note{})
	if result.Error != nil {
		return 0, custom_errors.GetDbError(result.Error, getType(item)+"->DeleteAll")
	}
	return db.RowsAffected, nil
}
