package db

import (
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
		mlog.Fatalf("new db error: %v", err.Error())
	}
	SetDB(mdb)
	if err := Migration(); err != nil {
		mlog.Fatalf("database migration error: %v", err.Error())
	}
	if conf.GetBool("mysql.debug", false) {
		db.Debug()
	}
	mlog.Info("init db success")
}

func Migration() error {
	DB := GetDB()
	if !DB.Migrator().HasTable(&v1.User{}) {
		if err := DB.Migrator().CreateTable(&v1.User{}); err != nil {
			mlog.Info("Error on migrate table user:", err.Error())
			return err
		}
	}
	if !DB.Migrator().HasTable(&v1.Reviewer{}) {
		if err := DB.Migrator().CreateTable(&v1.Reviewer{}); err != nil {
			mlog.Info("Error on migrate table reviewer:", err.Error())
			return err
		}
	}
	if !DB.Migrator().HasTable(&v1.Application{}) {
		if err := DB.Migrator().CreateTable(&v1.Application{}); err != nil {
			mlog.Info("Error on migrate table application:", err.Error())
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
