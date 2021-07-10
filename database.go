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
			id TEXT PRIMARY KEY DEFAULT CONCAT('mgtn_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			name TEXT NOT NULL UNIQUE,
			applied TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL
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
			CREATE TABLE accounts (
			    id              VARCHAR(40)  PRIMARY KEY DEFAULT CONCAT('acct_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			    created_at      TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
			    updated_at      TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
			    email           VARCHAR(512) NOT NULL UNIQUE,
				hashed_password VARCHAR(256) NOT NULL
			);

			CREATE TABLE websites (
			    id         VARCHAR(40)  PRIMARY KEY DEFAULT CONCAT('site_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			    account_id VARCHAR(40)  NOT NULL REFERENCES accounts(id),
			    created_at TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
			    updated_at TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
				domain     VARCHAR(256) NOT NULL UNIQUE
			);

			CREATE TABLE sessions (
			    id           VARCHAR(40) PRIMARY KEY DEFAULT CONCAT('sess_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			    account_id   VARCHAR(40) NOT NULL REFERENCES accounts(id),
			    refreshed_at TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
			    created_at   TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL
			);

			CREATE TABLE analytics_events (
			    id         VARCHAR(40)  PRIMARY KEY DEFAULT CONCAT('ae_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
				website_id VARCHAR(40)  NOT NULL REFERENCES websites(id),
			    created_at TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
			    name       VARCHAR(64)  NOT NULL
			);

			CREATE TABLE analytics_pageviews (
			    id           VARCHAR(40)  PRIMARY KEY DEFAULT CONCAT('ap_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
				website_id   VARCHAR(40)  NOT NULL REFERENCES websites(id),
			    sid          VARCHAR(64)  NOT NULL,
			    created_at   TIMESTAMP(3) WITH TIME ZONE DEFAULT NOW() NOT NULL,
				host         VARCHAR(512) NOT NULL,
				path         TEXT         NOT NULL,
				screen_size  VARCHAR(32)  NOT NULL,
				country_code VARCHAR(2)   NOT NULL
			);
		`)

		return err
	},
}}
