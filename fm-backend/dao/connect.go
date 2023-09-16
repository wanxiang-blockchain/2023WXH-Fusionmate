package dao

import (
	"fmt"
	"github.com/FusionMate/fm-backend/conf"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	fmdb *gorm.DB

	once sync.Once
)

func InitDB() {
	SetupDB()
	CreateCollectionTbl()
	CreateAIMessageTbl()
}

func SetupDB() {
	once.Do(connectDB)
}
func connectDB() {
	username := conf.GConfig.GetString("mysql.user")     //账号
	password := conf.GConfig.GetString("mysql.password") //密码
	host := conf.GConfig.GetString("mysql.host")         //数据库地址，可以是Ip或者域名:端口
	dbname := conf.GConfig.GetString("mysql.name")       //数据库名
	maxIdleConns := conf.GConfig.GetInt("mysql.maxIdleConns")
	maxOpenConns := conf.GConfig.GetInt("mysql.maxOpenConns")
	connMaxLifeTime := conf.GConfig.GetInt("mysql.connMaxLifeTime")

	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbname)

	//连接MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to open mysql:" + err.Error())
	}

	// 连通性测试
	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}

	// 连接
	sqldb.SetMaxIdleConns(maxIdleConns)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)
	fmdb = db
}

func CloseDB() error {
	db, err := fmdb.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
