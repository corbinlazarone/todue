package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/corbinlazarone/todue/cmd/internals/migrations"
	"github.com/corbinlazarone/todue/cmd/internals/models"
	"github.com/corbinlazarone/todue/cmd/internals/opencode"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // to load the pgx driver for database/sql
)

type application struct {
	infoLog        *log.Logger
	errLog         *log.Logger
	users          *models.UserModel
	courses        *models.CourseModel
	opencodeClient *opencode.Client
}

func main() {

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// db url is auto injected by docker-compose
	dbUrl := os.Getenv("DB_CONN")
	if dbUrl == "" {
		errLog.Fatal("DB_CONN env var not set")
	}

	dbPool, err := initDB(dbUrl)
	if err != nil {
		errLog.Fatal(err)
	}

	defer dbPool.Close()

	// create sql db connection for goose migrations
	sqlDB, err := sql.Open("pgx", dbUrl)
	if err != nil {
		errLog.Fatal(err)
	}
	defer sqlDB.Close()

	err = migrations.RunMigrations(sqlDB)
	if err != nil {
		errLog.Fatal("Failed to load migrations", err)
	}

	// create anthropic client
	apiKey := os.Getenv("OPENCODE_API_KEY")
	if apiKey == "" {
		errLog.Fatal("OPENCODE_API_KEY env var not set")
	}
	opencodeClient := opencode.NewClient(apiKey)

	app := &application{
		infoLog:        infoLog,
		errLog:         errLog,
		users:          &models.UserModel{DB: dbPool},
		courses:        &models.CourseModel{DB: dbPool},
		opencodeClient: opencodeClient,
	}

	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	infoLog.Println("Server running on port 4000...")
	err = srv.ListenAndServe()
	if err != nil {
		errLog.Fatal(err)
	}
}

func initDB(dataSource string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dataSource)

	if err != nil {
		return nil, err
	}

	// pinging the connection to see if it was successful
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
