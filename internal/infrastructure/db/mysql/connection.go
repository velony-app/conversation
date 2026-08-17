package mysql

import (
	"database/sql"
	"time"

	"github.com/velony-app/conversation/internal/conf"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(c *conf.Infrasturcture) (*sql.DB, error) {
	dsn := c.GetDatabase().GetSource()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
