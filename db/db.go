package db

import (
	"fmt"
	"time"

	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DataBase struct {
	db  *gorm.DB
	dsn string
}

var Db *DataBase

func InitDB(dsn_master, dsn_replicas1, dsn_replicas2 string) *DataBase {
	Db = &DataBase{
		dsn: dsn_master,
		db:  connectDB(dsn_master),
	}
	// init resolver

	Db.db.Use(
		dbresolver.Register(dbresolver.Config{
			// use `db2` as sources, `db3`, `db4` as replicas
			Sources:  []gorm.Dialector{mysql.Open(dsn_master)}, // if nil it will get default dsn_master when connected above
			Replicas: []gorm.Dialector{mysql.Open(dsn_replicas1)},
			// sources/replicas load balancing policy
			Policy: dbresolver.RandomPolicy{},
		}).Register(dbresolver.Config{
			// use `dsn_master` as sources (DB's default connection), `dsn_replicas2` as replicas for `User` reading
			Replicas: []gorm.Dialector{mysql.Open(dsn_replicas2)},
		}, &usermodels.User{}),
	)

	// config connection pool
	data, _ := Db.db.DB()
	data.SetConnMaxIdleTime(time.Hour)
	data.SetConnMaxLifetime(24 * time.Hour)
	data.SetMaxIdleConns(100)
	data.SetMaxOpenConns(200)

	// init transaction
	// dsta := Db.db.Clauses(dbresolver.Read).Begin()
	// dsta.Commit()
	// dsta.Where(nil, nil).Commit()
	return Db
}

func CloseDB() error {
	return nil
}

func connectDB(dsn string) *gorm.DB {
	var err error
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
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
