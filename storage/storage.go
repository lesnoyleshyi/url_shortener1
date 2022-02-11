package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Service interface {
	Save(string) (string, error)
	Retrieve(string) (string, error)
	Close() error
}

type Url struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type postgres struct{ conn *pgx.Conn }

func (p *postgres) Save(longUrl string) (string, error) {
	queryStr := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2);"

	ret, err := p.conn.Exec(context.Background(), queryStr, "dummy", longUrl)
	if err != nil || ret.RowsAffected() == 0 {
		return "", err
	}
	return ret.String(), nil
}

func (p *postgres) Retrieve(shortUrl string) (string, error) {
	queryStr := "SELECT * FROM urls WHERE short_url = $1;"

	ret, err := p.conn.Exec(context.Background(), queryStr, shortUrl)
	if err != nil || ret.RowsAffected() == 0 {
		return "", err
	}
	return ret.String(), nil
}

func (p *postgres) Close() error {
	return p.conn.Close(context.Background())
}

type DBconfig struct {
	User   string
	Passwd string
	Host   string
	Port   string
	DbName string
}

func NewConnection(dbconfig DBconfig) (Service, error) {
	connstr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbconfig.User, dbconfig.Passwd,
		dbconfig.Host, dbconfig.Port, dbconfig.DbName)
	createQuery := "CREATE TABLE IF NOT EXISTS urls (" +
		"short_url	varchar(10)		PRIMARY KEY, " +
		"long_url	varchar(2048)	NOT NULL UNIQUE" +
		");"

	connection, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting database: %v\n", err)
		os.Exit(1)
	}

	_, err = connection.Exec(context.Background(), createQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating table for url storage: %v\n", err)
		os.Exit(1)
	}
	return &postgres{conn: connection}, nil
}
