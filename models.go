package main

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	ID             string    `db:"id"         json:"id"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
	Email          string    `db:"email"      json:"email"`
	HashedPassword string    `db:"hashed_password" json:"-"`
}

type Event struct {
	ID        string    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func (s *dbStore) GetAccountByEmail(ctx context.Context, email string) (*Account, error) {
	var account Account
	err := s.db.GetContext(ctx, &account, `SELECT * FROM accounts WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *dbStore) CreateAccount(ctx context.Context, email, password string) (*Account, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var account Account
	err = s.db.QueryRowxContext(ctx, `
		INSERT INTO accounts (email, hashed_password) VALUES ($1, $2) 
		RETURNING *
	`, email, fmt.Sprintf("%s", hashed)).StructScan(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *dbStore) ListEvents(ctx context.Context) ([]Event, error) {
	var events []Event
	err := s.db.SelectContext(ctx, &events, `SELECT * FROM events ORDER BY created_at DESC`)
	return events, err
}

func (s *dbStore) CreateEvent(ctx context.Context) (*Event, error) {
	var event Event
	err := s.db.QueryRowxContext(ctx, `
		INSERT INTO events DEFAULT VALUES 
		RETURNING *
	`).StructScan(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
