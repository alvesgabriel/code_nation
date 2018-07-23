package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Actor struct {
	Actor string `json:"actor"`
	Quote string `json:"quote"`
}

//PORT port to be used
const PORT = "8080"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/v1/quote", quote()).Methods("GET", "OPTIONS")
	r.Handle("/v1/quote/{actor}", quoteByActor()).Methods("GET", "OPTIONS")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func quote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := getQuote()
		j, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)
		// w.WriteHeader(http.StatusNotImplemented)
	})
}

func quoteByActor() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		actor := strings.Replace(vars["actor"], "+", " ", -1)
		a := getQuoteActor(actor)
		if a.Actor == "" && a.Quote == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)
		// w.WriteHeader(http.StatusNotImplemented)
	})
}

func getQuote() Actor {
	var a Actor
	for a.Actor == "" && a.Quote == "" {
		query := `SELECT actor, detail
		FROM scripts
		ORDER BY RANDOM() LIMIT 1`
		row := db.QueryRow(query)
		row.Scan(&a.Actor, &a.Quote)
	}

	return a
}

func getQuoteActor(actor string) Actor {
	var a Actor
	query := `SELECT actor, detail
				FROM scripts
				WHERE actor = ?
				ORDER BY RANDOM() LIMIT 1`
	row := db.QueryRow(query, actor)
	row.Scan(&a.Actor, &a.Quote)

	return a
}
