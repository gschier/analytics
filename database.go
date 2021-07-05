package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type dbStore struct {
	db *sqlx.DB
}

func NewDB() *dbStore {
	return &dbStore{
		db: sqlx.MustConnect("postgres", Config.DatabaseURL),
	}
}

type Migration struct {
	Number  int
	Name    string
	Forward func(ctx context.Context, db *sqlx.DB) error
}

type dbRow struct {
	Id      string    `db:"id"`
	Name    string    `db:"name"`
	Applied time.Time `db:"applied"`
}

func mustMigrate(ctx context.Context, db *sqlx.DB) {
	err := migrate(ctx, db)
	if err != nil {
		panic(err)
	}
}

func migrate(ctx context.Context, db *sqlx.DB) error {
	fmt.Printf("[migrate] Running database migrations\n")

	// Create migrations table if it doesn't exist
	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS migrations (
			id TEXT PRIMARY KEY DEFAULT CONCAT('mgtn_', SUBSTRING(REPLACE(gen_random_uuid()::TEXT, '-', ''), 0, 16)),
			name TEXT NOT NULL UNIQUE,
			applied TIMESTAMP(3) DEFAULT NOW() NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	fmt.Printf("[migrate] Migrations table created\n")

	var completedMigrations []dbRow
	err = db.SelectContext(ctx, &completedMigrations, `
		SELECT * FROM migrations ORDER BY name ASC
	`)
	if err != nil {
		return err
	}

	for i, m := range migrations {
		name := fmt.Sprintf("%04d_%s", i+1, m.Name)
		hasRan := false
		for _, r := range completedMigrations {
			if r.Name == name {
				hasRan = true
				break
			}
		}

		if hasRan {
			fmt.Printf("[migrate] Skipping completed migration %s\n", name)
			continue
		}

		err := m.Forward(ctx, db)
		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, `INSERT INTO migrations (name) VALUES ($1)`, name)
		if err != nil {
			return err
		}

		fmt.Printf("[migrate] Ran migration %s\n", name)
		return nil
	}

	return nil
}

var migrations = []Migration{{
	Name: "create_tables",
	Forward: func(ctx context.Context, db *sqlx.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE OR REPLACE FUNCTION set_updated_at()   
			RETURNS TRIGGER AS $$
			BEGIN
				NEW.updated_at = NOW();
				RETURN NEW;   
			END;
			$$ LANGUAGE 'plpgsql';

			-- Create Accounts Table
			CREATE TABLE accounts (
			    id TEXT PRIMARY KEY DEFAULT CONCAT('acct_', SUBSTRING(REPLACE(gen_random_uuid()::TEXT, '-', ''), 0, 16)),
			    created_at TIMESTAMP(3) DEFAULT NOW() NOT NULL,
			    updated_at TIMESTAMP(3) DEFAULT NOW() NOT NULL,
			    email TEXT NOT NULL UNIQUE,
				hashed_password TEXT NOT NULL
			);
			CREATE TRIGGER set_account_updated BEFORE UPDATE ON accounts FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

			-- Create Websites Table
			CREATE TABLE websites (
			    id TEXT PRIMARY KEY DEFAULT CONCAT('site_', SUBSTRING(REPLACE(gen_random_uuid()::TEXT, '-', ''), 0, 16)),
			    created_at TIMESTAMP(3) DEFAULT NOW() NOT NULL,
			    updated_at TIMESTAMP(3) DEFAULT NOW() NOT NULL
			);
			CREATE TRIGGER set_websites_updated BEFORE UPDATE ON websites FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

			-- Create Events Table
			CREATE TABLE events (
			    id TEXT PRIMARY KEY DEFAULT CONCAT('evnt_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			    created_at TIMESTAMP(3) DEFAULT NOW() NOT NULL
			);
			CREATE TRIGGER set_events_updated BEFORE UPDATE ON events FOR EACH ROW EXECUTE PROCEDURE set_updated_at();
		`)

		return err
	},
}}
