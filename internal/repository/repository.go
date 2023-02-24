package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
