package repo

import (
	"database/sql"
	"effectivemobile-rest-service/internal/models"
	"effectivemobile-rest-service/internal/repo/services"
)

func RepositoryInstance(db *sql.DB) *Repository {
	return &Repository{
		RepositoryService: services.CreateServicesDB(db),
	}
}

type Repository struct {
	RepositoryService models.RepositoryService
}
