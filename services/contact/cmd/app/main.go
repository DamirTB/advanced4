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
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })
	// fmt.Println("Server listening on port 8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Server error:", err)
	// }

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
	//fmt.Println("Connected to the database successfully!")

	http.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from the index page!"))
	})
	
	fmt.Println("Server listening on port 4000")
	err = http.ListenAndServe("127.0.0.1:4000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
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
