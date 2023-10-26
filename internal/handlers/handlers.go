package handlers

import "effectivemobile-rest-service/internal/repo"

type Controllers struct {
	Repository *repo.Repository
}

func ControllersInstance(Repository *repo.Repository) *Controllers {
	return &Controllers{Repository: Repository}
}
