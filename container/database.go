package container

import (
	"errors"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	databaseMapping map[string]*gorm.DB
)

func init() {
	databaseMapping = map[string]*gorm.DB{}
}

func InitializeDatabase(name, dsn string) error {
	db, err := open(dsn)
	if err != nil {
		return err
	}

	databaseMapping[name] = db
	return nil
}

func open(connection string) (*gorm.DB, error) {
	if strings.HasPrefix(connection, "mysql://") {
		return gorm.Open(mysql.Open(connection[8:]), &gorm.Config{})
	}

	if strings.HasPrefix(connection, "sqlite://") {
		return gorm.Open(sqlite.Open(connection[9:]), &gorm.Config{})
	}

	return nil, errors.New("not support")
}

func GetDatabase(name string) *gorm.DB {
	return databaseMapping[name]
}
