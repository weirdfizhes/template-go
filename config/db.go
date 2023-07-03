package config

import (
	"fmt"
	"time"

	"template-go/tool/constants"
	"template-go/tool/logger"
	sqlTool "template-go/tool/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	uri := sqlTool.ConnString()
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		logger.LogFatalError(constants.ErrOpenConnection, err)
	}

	db.SetMaxIdleConns(constants.MaxIdleConn)
	db.SetMaxOpenConns(constants.MaxOpenConn)

	go pingDB(db)

	return db
}

func pingDB(db *sqlx.DB) {
	i := 1
	for range time.Tick(constants.TimePingInterval) {
		err := db.Ping()
		if err != nil {
			logger.LogFatalError(constants.ErrPingConnection, err)
		} else {
			logger.LogPrintSuccess(fmt.Sprintf("%s: Ping attempt %d", constants.SucPingConnection, i), nil)
		}

		i++
	}
}
