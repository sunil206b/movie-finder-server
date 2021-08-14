package main

import (
	"flag"
	"fmt"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"movie-finder/driver"
	"movie-finder/repository"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port  int
	env   string
	dbDSN string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config *config
	logger *log.Logger
	repo   *repository.Repo
}

func init() {
	gotenv.Load()
}

func main() {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Fatalf("failed to parse Elephant SQL URL %v\n", err)
	}

	cfg := new(config)
	flag.IntVar(&cfg.port, "port", 8085, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.dbDSN, "dsn", pgURL, "Postgres Connection URL")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := driver.ConnectPQSQL(pgURL)
	if err != nil {
		log.Fatalf("failed to connect Elephant SQL %v\n", err)
	}
	defer db.Close()
	app := &application{
		config: cfg,
		logger: logger,
		repo:   repository.NewRepo(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to run on port %d with ewrror %v\n", cfg.port, err)
	}
}
