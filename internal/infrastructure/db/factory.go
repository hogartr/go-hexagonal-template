package db

import (
	"database/sql"
	"fmt"

	"github.com/hogartr/go-hexagonal-template/internal/domain"
	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func NewUserRepoFromConfig(cfg config.Config) (domain.UserRepository, error) {
	dbConn, err := sql.Open(cfg.DBType, cfg.DBDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := Migrate(dbConn); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return NewUserRepo(dbConn), nil
}
