package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/williamzhou8818/gifts-miaoshia/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

// 建立数据库连接。代码讲解参见《双Token博客系统》
var (
	gifts_mysql      *gorm.DB
	gifts_mysql_once sync.Once
	dblog            ormlog.Interface

	gifts_redis     *redis.Client
	gift_redis_once sync.Once
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      ormlog.Silent,
			Colorful:      false,
		},
	)
}

func createMysqlDB(dbname, host, user, pass string, port int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true})
	if err != nil {
		util.LogRus.Panicf("connect to mysql use dsn %s failed: %s", dsn, err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxIdleConns(20)
	util.LogRus.Infof("connect to mysql db %s", dbname)
	return db
}

func GetGiftDBConnection() *gorm.DB {
	gifts_mysql_once.Do(func() {
		if gifts_mysql == nil {
			dbName := "gift"
			viper := util.CreateConfig("mysql")
			host := viper.GetString(dbName + ".host")
			port := viper.GetInt(dbName + ".port")
			user := viper.GetString(dbName + ".user")
			pass := viper.GetString(dbName + ".pass")
			gifts_mysql = createMysqlDB(dbName, host, user, pass, port)
		}
	})
	return gifts_mysql
}
