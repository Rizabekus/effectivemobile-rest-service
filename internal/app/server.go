package app

import (
	"database/sql"
	"effectivemobile-rest-service/internal/handlers"
	"effectivemobile-rest-service/internal/repo"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repo.RepositoryInstance(db)

	controller := handlers.ControllersInstance(repo)
	Routes(controller)
	defer db.Close()
}

func Routes(c *handlers.Controllers) {
	mux := http.NewServeMux()
	mux.HandleFunc("/people", c.People)
	fmt.Println("http://localhost:8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
