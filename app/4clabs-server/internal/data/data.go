package data

import (
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/pkg/qlogger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatProviderSet ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO warpped database client
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	db, err := gorm.Open(mysql.Open(c.Data.Database.Source), &gorm.Config{Logger: qlogger.NewDbTraceLogger(log)})
	if err != nil {
		return nil, nil, err
	}
	//
	//if err := db.Use(gormotel.Plugin); err != nil {
	//	return nil, nil, err
	//}

	d := &Data{
		DB: db,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		//if err := d.Rdb.Close(); err != nil {
		//	log.Error(err)
		//}
	}, nil
}
