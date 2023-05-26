package db

import (
	"fmt"
	"log"
	mlog "oa-review/logger"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     uint16
	DBName   string
}

func NewDBConfig(username, psw, host string, port uint16, dbname string) *DBConfig {
	return &DBConfig{
		Username: username,
		Password: psw,
		Host:     host,
		Port:     port,
		DBName:   dbname,
	}
}

func NewDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=true", config.Username, config.Password, config.Host, config.Port, config.DBName),
		DontSupportRenameColumn:  false,
		DontSupportRenameIndex:   false,
		DisableDatetimePrecision: false,
		DefaultStringSize:        256,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{SlowThreshold: 1 * time.Second, LogLevel: logger.Info},
		),
	})
	if err != nil {
		return nil, err
	}
	mlog.Info("new db successs")
	return db, nil
}
