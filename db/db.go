package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBase struct {
	db  *gorm.DB
	dsn string
}

var Db *DataBase

func InitDB(dsn string) *DataBase {
	Db = &DataBase{
		dsn: dsn,
		db:  connectDB(dsn),
	}
	return Db
}

func CloseDB() error {
	return nil
}

func connectDB(dsn string) *gorm.DB {
	var err error
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}
	return db
}

func (data *DataBase) NewSession() *gorm.DB {
	return data.db.Session(&gorm.Session{PrepareStmt: true})
}

func (data *DataBase) GetDB() *gorm.DB {
	return data.db
}

func (data *DataBase) AutoMigrate(dst ...interface{}) {
	if data.db != nil {
		data.db.AutoMigrate(dst...)
	}
}
