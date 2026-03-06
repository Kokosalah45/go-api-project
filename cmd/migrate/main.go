package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	m, err := newMigrate()
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

func RollbackMigrations() error {
	m, err := newMigrate()
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Migration rolled back successfully")
	return nil
}

func newMigrate() (*migrate.Migrate, error) {
	uri := buildMongoURI()

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		uri,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migrations: %w", err)
	}

	return m, nil
}

func buildMongoURI() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Print(host, port, username, password, dbName)
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		username, password, host, port, dbName,
	)
}

func main() {
	direction := "up"
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	switch direction {
	case "up":
		if err := RunMigrations(); err != nil {
			log.Fatalf("Migration up failed: %v", err)
		}
	case "down":
		if err := RollbackMigrations(); err != nil {
			log.Fatalf("Migration down failed: %v", err)
		}
	default:
		log.Fatalf("Unknown direction: %s (use 'up' or 'down')", direction)
	}
}
