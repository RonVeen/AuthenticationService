package internal

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const CreateTableUsers = `CREATE TABLE IF NOT EXISTS Users 
(
   uuid text,
   email text,
   status text,
   CONSTRAINT pk_users PRIMARY KEY(uuid) 
 );`

const CreateTableHash = `CREATE TABLE IF NOT EXISTS Secrets 
(
   uuid text,
   hash text,
   CONSTRAINT pk_secrets PRIMARY KEY(uuid) 
 );`

const CreateTableTokens = `CREATE TABLE IF NOT EXISTS Tokens
(
    uuid text,
    token text,
    expires timestamp,
    CONSTRAINT pk_tokens PRIMARY KEY(uuid)
);`

const CreateIndexTokens = `CREATE INDEX IF NOT EXISTS idx_token_token ON Tokens(token);`

func SetupDatabase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		"auth")
	conn, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	ensureDB(conn)

	return conn
}

func ensureDB(db *sqlx.DB) {
	var stmts = []string{CreateTableUsers, CreateTableHash, CreateTableTokens, CreateIndexTokens}
	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			log.Panic(err)
		}
	}
}
