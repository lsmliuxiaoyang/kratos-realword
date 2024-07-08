package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-realwd/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers. wire 需扫描的东西
var ProviderSet = wire.NewSet(NewData, NewDb, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{Db: db}, cleanup, nil
}

// NewDb 连接数据库
func NewDb(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	fmt.Println(c.Database.Source)
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}
