package repository

import (
	"ecommerce/common/log"
	"ecommerce/internal/sqlclient"
)

var DB sqlclient.ISqlClientConn

func InitRepo(db sqlclient.ISqlClientConn) {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	if err := db.GetDB().Ping(); err != nil {
		log.Fatal(err)
	}
	DB = db
}
