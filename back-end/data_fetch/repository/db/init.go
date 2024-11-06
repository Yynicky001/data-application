package db

import (
	"context"
	"data_fetch/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB

func init() {
	mysqlConf := config.Conf.Mysql
	dsn := strings.Join([]string{mysqlConf.UserName, ":", mysqlConf.Password,
		"@tcp(", mysqlConf.Host, ":", mysqlConf.Port, ")/",
		mysqlConf.Database, "?charset=", mysqlConf.Charset,
		"&parseTime=True&loc=Local"}, "")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,                         //连接数据库
		DefaultStringSize:         mysqlConf.DefaultStringSize, //string 类型字段的默认长度
		DisableDatetimePrecision:  true,                        //禁用 datetime precision, MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                        // 用 change 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                       // 不根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表名复数
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConn) //设置连接池
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConn) //最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxLifetime) * time.Second)
	_db = db
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
