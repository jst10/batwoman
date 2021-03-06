package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

func getType(item interface{}) string {
	if t := reflect.TypeOf(item); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

var db *gorm.DB

func InitDatabase() {
	fmt.Println("InitDatabase")
	dsn := "root:root@tcp(127.0.0.1:3306)/batwoman?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&NoteBody{})
	db.AutoMigrate(&Note{})
	db.AutoMigrate(&Folder{})
	db.AutoMigrate(&UserNote{})

}
