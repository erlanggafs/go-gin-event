package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	// Ganti koneksi ke MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3310)/gogin_event")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ganti instance untuk MySQL
	instance, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Buka file migrasi
	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi migrasi dengan MySQL
	m, err := migrate.NewWithInstance(
		"file", fSrc, "mysql", instance,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Eksekusi migrasi
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid direction. Use 'up' or 'down'")
	}
}
