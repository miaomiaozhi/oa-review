package db

import (
	"log"
	"oa-review/conf"
	mlog "oa-review/logger"
	v1 "oa-review/models/protoreq/v1"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDataBase(conf *conf.OaReviewConf) {
	mdb, err := NewDB(&DBConfig{
		Username: conf.MustGetString("mysql.username"),
		Password: conf.MustGetString("mysql.password"),
		Host:     conf.MustGetString("mysql.host"),
		Port:     uint16(conf.GetInt("mysql.port", 3306)),
		DBName:   conf.MustGetString("mysql.dbname"),
	})
	if err != nil {
		mlog.Fatal(err)
	}
	SetDB(mdb)
	if err := Migration(); err != nil {
		panic("migration error")
	}
	mlog.Info("init db success")
}

func Migration() error {
	DB := GetDB()
	if !DB.Migrator().HasTable(&v1.User{}) {
		if err := DB.Migrator().CreateTable(&v1.User{}); err != nil {
			log.Printf("Error on migrate table user: %v", err)
			return err
		}
	}
	if !DB.Migrator().HasTable(&v1.Reviewer{}) {
		if err := DB.Migrator().CreateTable(&v1.Reviewer{}); err != nil {
			log.Printf("Error on migrate table reviewer: %v", err)
			return err
		}
	}
	if !DB.Migrator().HasTable(&v1.Application{}) {
		if err := DB.Migrator().CreateTable(&v1.Application{}); err != nil {
			log.Printf("Error on migrate table Application: %v", err)
			return err
		}
	}
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		mlog.Fatal("init db must called before get db")
	}
	return db
}

func SetDB(mdb *gorm.DB) {
	db = mdb
}
