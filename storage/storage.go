package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Service interface {
	Save(string, string) (string, error)
	Retrieve(string, string) (string, error)
	Close() error
}

type postgres struct{ conn *pgx.Conn }

func (p *postgres) Save(shortUrl, longUrl string) (string, error) {
	queryStr := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2);"

	ret, err := p.conn.Exec(context.Background(), queryStr, shortUrl, longUrl)
	if err != nil || ret.RowsAffected() == 0 {
		return "", err
	}
	return shortUrl, nil
}

func (p *postgres) Retrieve(Url, ShortOrLong string) (string, error) {
	var queryStr string

	if ShortOrLong == "short" {
		queryStr = "SELECT short_url FROM urls WHERE long_url = $1;"
	} else if ShortOrLong == "long" {
		queryStr = "SELECT long_url FROM urls WHERE short_url = $1;"
	} else {
		return "", fmt.Errorf("invalid second argument for Retrieve() method")
	}
	var retUrl string
	err := p.conn.QueryRow(context.Background(), queryStr, Url).Scan(&retUrl)
	if err != nil {
		return "", err
	}
	return retUrl, nil
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

func (conf *DBconfig) ParseFromEnv() *DBconfig {
	conf.User = GetEnvOrDef("POSTGRES_USER", "go_user")
	conf.Passwd = GetEnvOrDef("POSTGRES_PASSWORD", "8246go")
	conf.Host = GetEnvOrDef("POSTGRES_HOST", "localhost")
	conf.Port = GetEnvOrDef("POSTGRES_PORT", "5432")
	conf.DbName = GetEnvOrDef("POSTGRES_DB", "url_storage")
	return conf
}

func GetEnvOrDef(env, defaultVal string) string {
	val, persist := os.LookupEnv(env)
	if !persist {
		return defaultVal
	}
	return val
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
