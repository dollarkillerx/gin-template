package simple

import (
	"github.com/dollarkillerx/gin-template/internal/conf"
	"github.com/dollarkillerx/gin-template/internal/pkg/models"
	"github.com/dollarkillerx/gin-template/internal/utils"
	"gorm.io/gorm"

	"sync"
)

type Simple struct {
	db *gorm.DB

	inventoryMu sync.Mutex
}

func NewSimple(conf *conf.PgSQLConfig) (*Simple, error) {
	sql, err := utils.InitPgSQL(conf)
	if err != nil {
		return nil, err
	}

	sql.AutoMigrate(
		&models.User{},
	)

	return &Simple{
		db: sql,
	}, nil
}

func (s *Simple) DB() *gorm.DB {
	return s.db
}
