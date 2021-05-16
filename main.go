package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	sqlFunctions "url-shortener-mysql/sql"

	"github.com/gorilla/mux"
)

func main() {
	db, err := sqlFunctions.OpenConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	dbHandler := newDatabaseHandler(db)
	r := mainRouter(db, dbHandler)

	http.ListenAndServe(":8000", r)
}

func mainRouter(db *sql.DB, fallback http.HandlerFunc) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		err := sqlFunctions.InsertPath(db, query.Get("path"), query.Get("url"))
		if err != nil {
			json.NewEncoder(w).Encode("Insertion Failed!")
			return
		}
		json.NewEncoder(w).Encode("Successfully Inserted!")
	})
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	})
	return router
}

func newDatabaseHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathFromURL, err := sqlFunctions.GetAllPath(db)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			path := r.URL.Path
			if url, ok := pathFromURL[path]; ok {
				http.Redirect(w, r, url, http.StatusFound)
				return
			}
		}
		json.NewEncoder(w).Encode("Not Found!")
	}
}
