package db

import (
	"ecloudsystem/configs"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Client *dbRepo

type Repo interface {
	GetDbW() *gorm.DB
	DbWClose() error
	//GetDbR() *gorm.DB
	//DbRClose() error
}

type dbRepo struct {
	DbW *gorm.DB
	//DbR *gorm.DB
}

func InitDb() (Repo, error) {
	cfg := configs.Get().MySQL
	dbw, err := dbConnect(cfg.Db.User, cfg.Db.Pass, cfg.Db.Addr, cfg.Db.Name)
	if err != nil {
		return nil, err
	}

	//dbw, err := dbConnect(cfg.Write.User, cfg.Write.Pass, cfg.Write.Addr, cfg.Write.Name)
	//if err != nil {
	//	return nil, err
	//}

	Client = &dbRepo{
		//DbR: dbr,
		DbW: dbw,
	}
	return Client, nil
}


//func (d *dbRepo) GetDbR() *gorm.DB {
//	return d.DbR
//}

func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

//func (d *dbRepo) DbRClose() error {
//	sqlDB, err := d.DbW.DB()
//	if err != nil {
//		return err
//	}
//	return sqlDB.Close()
//}

func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	cfg := configs.Get().MySQL.Gorm

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)

	// 使用插件
	//db.Use(&TracePlugin{})

	return db, nil
}