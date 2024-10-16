package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"github.com/tursodatabase/go-libsql"
)

type DB struct {
	*Queries
	db *sql.DB
}

func NewDBWrapper(db *sql.DB) *DB {
	return &DB{
		Queries: New(db),
		db:      db,
	}
}

// func Open(dataSourceName string) (*DB, error) {
// 	db, err := sql.Open("libsql", dataSourceName)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening database: %w", err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("error connecting to database: %w", err)
// 	}

// 	return NewDBWrapper(db), nil
// }

func Open(dataSourceName string) (*DB, error) {
	godotenv.Load()

	envAuthToken := os.Getenv("AUTH_TOKEN")
	envURL := os.Getenv("LIBSQL_URL")

	var authToken string
	var primaryUrl string

	if envURL != "" {
		primaryUrl = envURL
	}

	if envAuthToken != "" {
		authToken = envAuthToken
	}

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dataSourceName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
	)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}
	defer connector.Close()

	db := sql.OpenDB(connector)

	return NewDBWrapper(db), nil
}

func (w *DB) Close() error {
	return w.db.Close()
}

func (w *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return w.db.BeginTx(ctx, opts)
}

func (w *DB) WithTx(tx *sql.Tx) *Queries {
	return New(tx)
}
