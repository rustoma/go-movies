package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
}
