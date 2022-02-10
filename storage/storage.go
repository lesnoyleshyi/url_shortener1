package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Url struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type DBconfig struct {
	User   string
	Passwd string
	Host   string
	Port   string
	DbName string
}

func NewConnection(dbconfig DBconfig) (*pgx.Conn, error) {
	connstr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbconfig.User, dbconfig.Passwd,
		dbconfig.Host, dbconfig.Port, dbconfig.DbName)
	createQuery := "CREATE TABLE IF NOT EXISTS urls (" +
		"short_url	varchar(10)		PRIMARY KEY, " +
		"long_url	varchar(2048)	NOT NULL UNIQUE" +
		");"

	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), createQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating table for url storage: %v\n", err)
		os.Exit(1)
	}
	return conn, nil
}
