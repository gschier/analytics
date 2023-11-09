package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

var _db *sqlx.DB

func GetDB() *sqlx.DB {
	if _db != nil {
		return _db
	}

	var err error
	for i := 0; i < 5; i++ {
		_db, err = sqlx.Connect("postgres", Config.DatabaseURL)
		if err != nil {
			logger.Warn("Failed to connect to database", "error", err)
			time.Sleep(1 * time.Second)
		} else {
			return _db
		}
	}

	panic("Failed to connect to database")
}

type Migration struct {
	Name    string
	Forward func(ctx context.Context, db *sqlx.DB) error
}

type dbRow struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	AppliedAt time.Time `db:"applied_at"`
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
			id         TEXT PRIMARY KEY DEFAULT CONCAT('mgtn_', REPLACE(gen_random_uuid()::TEXT, '-', '')),
			name       TEXT NOT NULL UNIQUE,
			applied_at TIMESTAMP(3) WITH TIME ZONE
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

		_, err = db.NamedExecContext(ctx, `
			INSERT INTO migrations (name, applied_at) 
			VALUES (:name, :applied_at)
		`, &dbRow{
			Name:      name,
			AppliedAt: time.Now(),
		})
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
			    id              VARCHAR(40)  PRIMARY KEY ,
			    created_at      TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW(),
			    updated_at      TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW(),
			    email           VARCHAR(512) NOT NULL UNIQUE,
				hashed_password VARCHAR(256) NOT NULL
			);

			CREATE TABLE websites (
			    id         VARCHAR(40)  PRIMARY KEY,
			    account_id VARCHAR(40)  NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			    created_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW(),
			    updated_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW(),
				domain     VARCHAR(256) NOT NULL UNIQUE
			);

			CREATE TABLE sessions (
			    id           VARCHAR(40) PRIMARY KEY,
			    account_id   VARCHAR(40) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			    refreshed_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW(),
			    created_at   TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT NOW()
			);

			-- Analytics tables don't have unique constraints, FKs, or PKs for fast inserts

			CREATE TABLE analytics_events (
			    id         VARCHAR(64)  NOT NULL,
			    sid        VARCHAR(64)  NOT NULL,
				website_id VARCHAR(40)  NOT NULL,
			    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
			    name       VARCHAR(64)  NOT NULL
			);
			CREATE INDEX analytics_events__website_id_created_at ON analytics_events (website_id, created_at);

			CREATE TABLE analytics_pageviews (
			    id           VARCHAR(64)  NOT NULL,
			    sid          VARCHAR(64)  NOT NULL,
				website_id   VARCHAR(40)  NOT NULL,
			    created_at   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
				host         VARCHAR(512) NOT NULL,
				path         TEXT         NOT NULL,
				screen_size  VARCHAR(32)  NOT NULL,
				country_code VARCHAR(2)   NOT NULL,
				user_agent   TEXT         NOT NULL
			);
			CREATE INDEX analytics_pageviews__website_id_created_at ON analytics_pageviews (website_id, created_at);
		`)

		return err
	},
}, {
	Name: "add_event_attributes",
	Forward: func(ctx context.Context, db *sqlx.DB) error {
		_, err := db.ExecContext(ctx, `
			ALTER TABLE analytics_events ADD COLUMN attributes JSONB NOT NULL DEFAULT '{}';
		`)

		return err
	},
}, {
	Name: "add_more_event_properties",
	Forward: func(ctx context.Context, db *sqlx.DB) error {
		_, err := db.ExecContext(ctx, `
			ALTER TABLE analytics_events 
				ADD COLUMN screen_size    VARCHAR(32)    NOT NULL DEFAULT '',
				ADD COLUMN country_code   VARCHAR(2)     NOT NULL DEFAULT '',
				ADD COLUMN version        VARCHAR(32)    NOT NULL DEFAULT '',
				ADD COLUMN platform       VARCHAR(16)    NOT NULL DEFAULT ''
		`)

		return err
	},
}}
