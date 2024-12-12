package repositories

import (
	"fmt"
	"net/url"
	"os"

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
		return nil, fmt.Errorf("failed to open connection to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return db, nil
}

func connectString(cfg Config) string {

	// databaseURL := &url.URL{
	// 	Scheme: "postgres",
	// 	Host:   cfg.Host + ":" + cfg.Port,
	// 	Path:   "/" + cfg.DBName,
	// 	User:   url.UserPassword("dev", "dev1234"),
	// }
	databaseURL := &url.URL{
		Scheme: "postgres",
		Host:   os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PORT"),
		Path:   "/" + os.Getenv("POSTGRES_DB"),
		User:   url.UserPassword(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")),
	}
	rawQuery := make(url.Values)
	rawQuery.Set("ssl_mode", "disable")
	databaseURL.RawQuery = rawQuery.Encode()
	return databaseURL.String()
}
