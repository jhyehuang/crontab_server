package middleware

import (
	"fmt"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type GormDB struct {
}

func NewDbConfig() *GormDB {
	return &GormDB{}
}

func (this *GormDB) GormDB() *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			configs.GetValue(configs.MysqlUsername, configs.Get().MySQL.Username),
			configs.GetValue(configs.MysqlPassword, configs.Get().MySQL.Password),
			configs.GetValue(configs.MysqlHost, configs.Get().MySQL.Host),
			configs.GetValue(configs.MysqlPort, configs.Get().MySQL.Port),
			configs.GetValue(configs.MysqlDatabase, configs.Get().MySQL.Db),
			"Asia%2FShanghai"),
		//DefaultStringSize:         256,   // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: logger.Default.LogMode(logger.Silent), // 这里设置为 Silent 表示关闭 GORM 的日志输出
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("database connected success...")
	sqlDB, err := db.DB()
	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(configs.Get().MySQL.MaxIdleConns)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(configs.Get().MySQL.MaxOpenConns)
	//设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	models.SetupDatabase(db)
	return db
}
