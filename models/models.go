package models

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Model is basic internal information
type Model struct {
	db     *pgxpool.Pool //nolint
	logger *slog.Logger  //nolint
}

// SetConfig adds database and logger instances for handling stuff
func (m *Model) SetConfig(db *pgxpool.Pool, logger *slog.Logger) {
	m.db = db
	m.logger = logger
}
