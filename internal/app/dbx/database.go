package dbx

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/internal/app/gormx"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
	"os"
	"path/filepath"
)

var gormDBM map[string]*gorm.DB
//var sqlxDBM map[string]*sqlx.DB


// InitGormDB 初始化gorm存储
func InitGormDB() (map[string]func(), error) {
	gormDBM = make(map[string]*gorm.DB, 10)
	cleanFuncMap := make(map[string]func(), 10)
	c := config.C.Database.Gorm
	connections := config.C.Database.Connections
	log.Println(connections)
	for _, conn := range connections {
		if conn.Orm == "gorm" {
			db, cleanFunc, err := NewGormDB(&c, &conn)
			if err != nil {
				return cleanFuncMap, err
			}
			gormDBM[conn.Tag] = db
			cleanFuncMap[conn.Tag] = cleanFunc
		} else if conn.Orm == "sqlx" {

		}
	}
	return cleanFuncMap, nil
}

func NewGormDB(c *config.Gorm, conn *config.Connection)(*gorm.DB, func(), error) {
	return gormx.NewDB(&gormx.Config{
		Debug:        c.Debug,
		DBType:       conn.Driver,
		DSN:          DSN(conn),
		MaxIdleConns: c.MaxIdleConns,
		MaxLifetime:  c.MaxLifetime,
		MaxOpenConns: c.MaxOpenConns,
		TablePrefix:  c.TablePrefix,
	})
}

func NewSqlxDB() {

}

func DSN(c *config.Connection) string {
	var dsn string
	if c.Driver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			c.Username, c.Password, c.Host, c.Port, c.Database, c.Options)
	} else if c.Driver == "sqlite3" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			c.Username, c.Password, c.Host, c.Port, c.Database, c.Options)
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	} else if c.Driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			c.Host, c.Port, c.Username, c.Database, c.Password)
	} else {
		log.Println(c.Driver)
		logger.Error("Unsupported Database")
	}
	return dsn
}

func DefaultDSN() string {
	connections := config.C.Database.Connections
	for _, conn := range connections {
		if conn.Tag == "default" {
			return DSN(&conn)
		}
	}
	logger.Error("Default database connection config does not exist!")
	return ""
}

func GormDB(tag string) *gorm.DB  {
	db, ok := gormDBM[tag]; if !ok {
		logger.Error("GormDB instance does not exist!")
	}
	return db
}