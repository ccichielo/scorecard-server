package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Score struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
}

func AddScore(w http.ResponseWriter, r *http.Request) {
	var score Score
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connectionString := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO scorecard.scores(id, name, score) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("Prepared Statement Failure")
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(score.Id, score.Name, score.Score)
	if err != nil {
		fmt.Println("Exec Failure")
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
