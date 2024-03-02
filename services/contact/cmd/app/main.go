package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "log"
	"net/http"
	"os"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "godb"
	)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := openDB(connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
	defer db.Close()
	http.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from the index page!"))
	})
	fmt.Println("Server listening on port 4000")
	err = http.ListenAndServe("127.0.0.1:4000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func openDB(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
