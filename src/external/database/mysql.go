package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "embed"

	"github.com/go-sql-driver/mysql"
	"github.com/x-syaifullah-x/go-crud/src/internal/config"
	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

var db *sql.DB
var once sync.Once

//go:embed query/01_create_table_users.sql
var queryCreateTableUsers []byte

//go:embed query/01_create_trigger_before_insert_table_users.sql
var queryCreateTriggerBeforeInsertTableUsers []byte

//go:embed query/01_create_trigger_before_update_table_users.sql
var queryCreateTriggerBeforeUpdateTableUsers []byte

func Instance() *sql.DB {
	once.Do(func() {
		config := config.GetConfig()
		cfg := mysql.Config{
			User:                 config.DB.User,
			Passwd:               config.DB.Password,
			Net:                  "tcp",
			Addr:                 config.DB.Host,
			DBName:               config.DB.Name,
			AllowNativePasswords: true,
		}

		var err error

		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			logger.Fatal(err.Error())
		}

		err = db.Ping()
		if err != nil {
			logger.Fatal(err.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			logger.Fatal(err.Error())
		}
		defer tx.Commit()

		var rows *sql.Row
		rows = tx.QueryRow(string(queryCreateTableUsers))
		if rows.Err() != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Fatal(rows.Err())
		}

		rows = tx.QueryRow(string(queryCreateTriggerBeforeInsertTableUsers))
		if rows.Err() != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Fatal(rows.Err())
		}

		rows = tx.QueryRow(string(queryCreateTriggerBeforeUpdateTableUsers))
		if rows.Err() != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Fatal(rows.Err())
		}

		// query := fmt.Sprintf(
		// 	"CREATE INDEX IF NOT EXISTS %1s ON %2s(%3s)",
		// 	"name",
		// 	"users",
		// 	"name",
		// )
		// rows = tx.QueryRow(query)
		// if rows.Err() != nil {
		// 	err := tx.Rollback()
		// 	if err != nil {
		// 		logger.Fatal(err)
		// 	}
		// 	logger.Fatal(rows.Err())
		// }

		// query = fmt.Sprintf(
		// 	"CREATE INDEX IF NOT EXISTS %1s ON %2s(%3s)",
		// 	"username",
		// 	"users",
		// 	"username",
		// )
		// rows = tx.QueryRow(query)
		// if rows.Err() != nil {
		// 	err := tx.Rollback()
		// 	if err != nil {
		// 		logger.Fatal(err)
		// 	}
		// 	logger.Fatal(rows.Err())
		// }

		query := fmt.Sprintf(
			"CREATE VIEW IF NOT EXISTS %1s AS SELECT * FROM %2s WHERE %3s = 0",
			"users_active",
			"users",
			"delete_at",
		)
		rows = tx.QueryRow(query)
		if rows.Err() != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Fatal(rows.Err())
		}

		db.SetConnMaxIdleTime(time.Duration(config.DB.ConnectionPool.MaxIdletimeConnection) * time.Second)
		db.SetConnMaxLifetime(time.Duration(config.DB.ConnectionPool.MaxLifetimeConnection) * time.Second)
		db.SetMaxOpenConns(int(config.DB.ConnectionPool.MaxOpenConnetcion))
		db.SetMaxIdleConns(int(config.DB.ConnectionPool.MaxIdleConnection))
	})
	return db
}
