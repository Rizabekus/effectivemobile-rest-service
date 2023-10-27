package handlers

import (
	"effectivemobile-rest-service/internal/repo"
	"net/http"
)

type Controllers struct {
	Repository *repo.Repository
}

func ControllersInstance(Repository *repo.Repository) *Controllers {
	return &Controllers{Repository: Repository}
}

func (c *Controllers) People(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch len(r.URL.Query()) {
		case 0:
		case 1:
		default:

		}
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
