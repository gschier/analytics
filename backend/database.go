package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

var _db *sqlx.DB

func GetDB() DBLike {
	if _db == nil {
		_db = sqlx.MustConnect("postgres", Config.DatabaseURL)
	}
	return _db
}

type Migration struct {
	Number  int
	Name    string
	Forward func(ctx context.Context, db DBLike) error
}

type dbRow struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	AppliedAt time.Time `db:"applied_at"`
}

func mustMigrate(ctx context.Context, db DBLike) {
	err := migrate(ctx, db)
	if err != nil {
		panic(err)
	}
}

func migrate(ctx context.Context, db DBLike) error {
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
	Forward: func(ctx context.Context, db DBLike) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE accounts (
			    id              VARCHAR(40)  PRIMARY KEY ,
			    created_at      TIMESTAMP(3) WITH TIME ZONE,
			    updated_at      TIMESTAMP(3) WITH TIME ZONE,
			    email           VARCHAR(512) NOT NULL UNIQUE,
				hashed_password VARCHAR(256) NOT NULL
			);

			CREATE TABLE websites (
			    id         VARCHAR(40)  PRIMARY KEY,
			    account_id VARCHAR(40)  NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			    created_at TIMESTAMP(3) WITH TIME ZONE,
			    updated_at TIMESTAMP(3) WITH TIME ZONE,
				domain     VARCHAR(256) NOT NULL UNIQUE
			);

			CREATE TABLE sessions (
			    id           VARCHAR(40) PRIMARY KEY,
			    account_id   VARCHAR(40) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			    refreshed_at TIMESTAMP(3) WITH TIME ZONE,
			    created_at   TIMESTAMP(3) WITH TIME ZONE
			);

			CREATE TABLE analytics_events (
			    id         VARCHAR(40)  PRIMARY KEY,
				website_id VARCHAR(40)  NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
			    created_at TIMESTAMP(3) WITH TIME ZONE,
			    sid        VARCHAR(64)  NOT NULL,
			    name       VARCHAR(64)  NOT NULL
			);

			CREATE TABLE analytics_pageviews (
			    id           VARCHAR(40)  PRIMARY KEY,
				website_id   VARCHAR(40)  NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
			    sid          VARCHAR(64)  NOT NULL,
			    created_at   TIMESTAMP(3) WITH TIME ZONE,
				host         VARCHAR(512) NOT NULL,
				path         TEXT         NOT NULL,
				screen_size  VARCHAR(32)  NOT NULL,
				country_code VARCHAR(2)   NOT NULL,
				user_agent   TEXT         NOT NULL
			);
		`)

		return err
	},
}}