package dao

// var DB *gorm.DB

//func init() {
//	err := InitDataBase()
//	if err != nil {
//		panic(err)
//	}
//}
//
//func getDns() string {
//	// "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
//	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		conf.Username, conf.Password, conf.Host, conf.Port, conf.Dbname)
//	return dns
//}
//
//func InitDataBase() error {
//	dns := getDns()
//	log.Println("Init database now, DNS is", dns)
//
//	db, err := gorm.Open(mysql.New(mysql.Config{
//		DSN:                       dns,   // DSN data source name
//		DefaultStringSize:         256,   // string 类型字段的默认长度
//		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
//		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
//		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
//		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
//	}), &gorm.Config{})
//
//	if err != nil {
//		log.Printf("Error on open database:%v\n", err)
//		return err
//	}
//	sqlDB, err := db.DB()
//
//	if err != nil {
//		log.Printf("Error on open database:%v\n", err)
//		return err
//	}
//	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
//	sqlDB.SetMaxIdleConns(10)
//
//	// SetMaxOpenConns sets the maximum number of open connections to the database.
//	sqlDB.SetMaxOpenConns(100)
//
//	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//	sqlDB.SetConnMaxLifetime(time.Hour)
//	DB = db
//
//	// migrate
//	if err := migration(); err != nil {
//		return err
//	}
//
//	log.Println("Init database successfully")
//	return nil
//}
//
// func migration() error {
// 	if !DB.Migrator().HasTable(&v1.User{}) {
// 		if err := DB.Migrator().CreateTable(&v1.User{}); err != nil {
// 			log.Printf("Error on migrate table user: %v", err)
// 			return err
// 		}
// 	}
// 	if !DB.Migrator().HasTable(&v1.Reviewer{}) {
// 		if err := DB.Migrator().CreateTable(&v1.Reviewer{}); err != nil {
// 			log.Printf("Error on migrate table reviewer: %v", err)
// 			return err
// 		}
// 	}
// 	if !DB.Migrator().HasTable(&v1.Application{}) {
// 		if err := DB.Migrator().CreateTable(&v1.Application{}); err != nil {
// 			log.Printf("Error on migrate table Application: %v", err)
// 			return err
// 		}
// 	}
// 	return nil
// }
