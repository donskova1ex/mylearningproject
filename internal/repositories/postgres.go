package repositories

import (
	"net/url"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectString(cfg))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectString(cfg Config) string {

	databaseURL := &url.URL{
		Scheme: "postgres",
		Host:   cfg.Host + ":" + cfg.Port,
		Path:   "/" + cfg.DBName,
		User:   url.UserPassword(cfg.Username, cfg.Password),
	}
	rawQuery := make(url.Values)
	rawQuery.Set("ssl_mode", "disable")
	databaseURL.RawQuery = rawQuery.Encode()
	return databaseURL.String()
}
