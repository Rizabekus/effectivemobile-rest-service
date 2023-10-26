package app

import (
	"database/sql"
	"effectivemobile-rest-service/internal/repo"
	"log"
)

func Run() {
	db, err := sql.Open("sqlite3", "./sql/database.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repo.RepositoryInstance(db)

	controller := controllers.ControllersInstance(repo)
	Routes(controller)
	defer db.Close()
}

func Routes(c controller) {
}
