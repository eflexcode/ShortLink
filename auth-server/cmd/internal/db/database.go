package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type ConfigDb struct {
	DbType       string
	Addr         string
	MaxOpenConn  int
	MaxIdealConn int
	MaxIdealTime string
}

func DatabaseConnect(c ConfigDb) (s *sql.DB, e error) {

	db, err := sql.Open(c.DbType, c.Addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.MaxIdealConn)
	db.SetMaxOpenConns(c.MaxOpenConn)
	db.SetConnMaxIdleTime(time.Duration(c.MaxIdealConn))

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	
	return db,nil

}

