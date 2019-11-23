package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init(){
	var (
		err error
		dbType, dbName, user, password, host string
	)
	dbType = "mysql"
	dbName = "yidui_test"
	user = "root"
	password = "root"
	host = "127.0.0.1:3306"
	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println("err", err)
	}

	Db = db
	err = db.DB().Ping()
	if err != nil {
		log.Println("ping", err)
	}
}

func CloseDB(){
	defer Db.Close()
}
