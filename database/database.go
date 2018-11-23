package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spring2go/gravitee/config"

	// Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

// NewDatabase returns a gorm.DB struct, gorm.DB.DB() returns a database handle
// see http://golang.org/pkg/database/sql/#DB
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	// Mysql
	if cfg.Database.Type == "mysql" {
		// Connection args
		// see https://github.com/go-sql-driver/mysql#dsn-data-source-name
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DatabaseName,
		)

		db, err := gorm.Open(cfg.Database.Type, args)
		if err != nil {
			return db, err
		}

		// Max idle connections
		db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)

		// Database logging
		db.LogMode(cfg.IsDevelopment)

		return db, nil
	}

	// Database type not supported
	return nil, fmt.Errorf("Database type %s not supported", cfg.Database.Type)
}
