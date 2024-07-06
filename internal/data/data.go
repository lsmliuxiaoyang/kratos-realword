package data

import (
	"gorm.io/gorm"
	"kratos-demo/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gorm.io/driver/mysql"
	"gorm.io/plugin/dbresolver"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data 访问第三方的数据，缓存等。面向存储做一些逻辑
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

type ApolloDb *gorm.DB

func NewMysql(conf *conf.Data, logger log.Logger) ApolloDb {
	logDb := log.NewHelper(log.With(logger, "x_module", "data/gorm-opc"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		logDb.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		logDb.Fatalf("get portrait to mysql: %v", err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)

	err = db.Use(
		dbresolver.Register(dbresolver.Config{Replicas: []gorm.Dialector{mysql.Open(conf.Database.Source)}}).
			SetConnMaxLifetime(time.Hour),
	)
	if err != nil {
		logDb.Fatalf("failed db use to mysql: %v", err)
	}
	logDb.Info("init apollo db")
	return db
}
