package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql" // register driver
)

func main() {
	// TODO: Replace with your actual Railway app domain.
	const appDomain = "libsql-production-b53d.up.railway.app"

	// TODO: generate your own authToken (see README)
	const authToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MDYwMTg5NTQsInN1YiI6InVzZXIifQ.GaOet8Jv_WOUYEv71oSp-ps0XfJAzAkQY34W3zxAcE70xFZa8KRGpkt4ql-7IDXBHwVWpckHnynMgthsBpghDg" // Example only — don’t hardcode secrets in production.

	// Construct the full DSN for connecting to LibSQL.
	dsn := fmt.Sprintf("libsql://%s?authToken=%s", appDomain, authToken)

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
