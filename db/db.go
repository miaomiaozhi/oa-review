package db

import (
	bean "oa-review/bean"
	"oa-review/conf"
	mlog "oa-review/logger"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDataBase(conf *conf.OaReviewConf) {
	dbConfig := &DBConfig{
		Username: conf.MustGetString("mysql.username"),
		Password: conf.MustGetString("mysql.password"),
		Host:     conf.MustGetString("mysql.host"),
		Port:     uint16(conf.GetInt("mysql.port", 3306)),
		DBName:   conf.MustGetString("mysql.dbname"),
	}
	mdb, err := NewDB(dbConfig)
	// mlog.Debug("db config info is")
	// mlog.Debug(dbConfig.Username)
	// mlog.Debug(dbConfig.Password)
	// mlog.Debug(dbConfig.Host)
	// mlog.Debug(dbConfig.Port)
	// mlog.Debug(dbConfig.DBName)

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
	if !DB.Migrator().HasTable(&bean.User{}) {
		if err := DB.Migrator().CreateTable(&bean.User{}); err != nil {
			mlog.Info("Error on migrate table user:", err.Error())
			return err
		}
	}
	if !DB.Migrator().HasTable(&bean.Reviewer{}) {
		if err := DB.Migrator().CreateTable(&bean.Reviewer{}); err != nil {
			mlog.Info("Error on migrate table reviewer:", err.Error())
			return err
		}
	}
	if !DB.Migrator().HasTable(&bean.Application{}) {
		if err := DB.Migrator().CreateTable(&bean.Application{}); err != nil {
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
