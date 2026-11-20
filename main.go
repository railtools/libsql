package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql" // register driver
)

func main() {
	// Your Railway public domain, e.g. "libsql-production-xxxx.up.railway.app".
	// A full URL such as "http://localhost:8080" also works.
	appDomain := os.Getenv("LIBSQL_DOMAIN")
	// JWT generated with `go run ./cmd/jwtgen` (see README).
	authToken := os.Getenv("LIBSQL_AUTH_TOKEN")
	if appDomain == "" || authToken == "" {
		log.Fatal("set LIBSQL_DOMAIN and LIBSQL_AUTH_TOKEN (see README)")
	}

	dsn := appDomain
	if !strings.Contains(dsn, "://") {
		dsn = "libsql://" + dsn
	}
	dsn = fmt.Sprintf("%s?authToken=%s", dsn, authToken)

	db, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatalf("open failed: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	_, err = db.ExecContext(ctx, `DROP TABLE IF EXISTS users`)
	if err != nil {
		log.Fatalf("drop failed: %v", err)
	}

	_, err = db.ExecContext(
		ctx, `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT
		)
	`,
	)
	if err != nil {
		log.Fatalf("create failed: %v", err)
	}

	_, _ = db.ExecContext(ctx, `INSERT INTO users (name) VALUES ('Alice')`)
	_, _ = db.ExecContext(ctx, `INSERT INTO users (name) VALUES ('Bob')`)

	rows, err := db.QueryContext(ctx, `SELECT id, name FROM users ORDER BY id`)
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatalf("scan failed: %v", err)
		}
		fmt.Printf("User %d: %s\n", id, name)
	}
}
