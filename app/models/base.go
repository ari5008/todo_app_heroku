package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
	}

	conn, err := pq.ParseURL(url)
	if err != nil {
			log.Fatalf("Failed to parse DATABASE_URL: %s", err)
	}

	conn += " sslmode=require1"

	Db, err = sql.Open("postgres", conn)
	if err != nil {
			log.Fatalf("Failed to connect to database: %s", err)
	}

	if err := Db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %s", err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
