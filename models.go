package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	ID             string    `db:"id"              json:"id"`
	CreatedAt      time.Time `db:"created_at"      json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at"      json:"updatedAt"`
	Email          string    `db:"email"           json:"email"`
	HashedPassword string    `db:"hashed_password" json:"-"`
}

type Session struct {
	ID          string    `db:"id"           json:"id"`
	AccountID   string    `db:"account_id"   json:"accountId"`
	CreatedAt   time.Time `db:"created_at"   json:"createdAt"`
	RefreshedAt time.Time `db:"refreshed_at" json:"refreshedAt"`
}

type Website struct {
	ID        string    `db:"id"         json:"id"`
	AccountId string    `db:"account_id" json:"accountId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Domain    string    `db:"domain"     json:"domain"`
}

type AnalyticsEvent struct {
	ID         string    `db:"id"          json:"id"`
	WebsiteID  string    `db:"website_id"  json:"websiteId"`
	CreatedAt  time.Time `db:"created_at"  json:"createdAt"`
	SessionKey string    `db:"session_key" json:"sessionKey"`
	Name       string    `db:"name"        json:"name"`
}

type AnalyticsPageview struct {
	ID          string    `db:"id"           json:"id"`
	WebsiteID   string    `db:"website_id"   json:"websiteId"`
	SID         string    `db:"sid"          json:"sid"`
	CreatedAt   time.Time `db:"created_at"   json:"createdAt"`
	Host        string    `db:"host"         json:"host"`
	Path        string    `db:"path"         json:"path"`
	ScreenSize  string    `db:"screen_size"  json:"screenSize"`
	CountryCode string    `db:"country_code" json:"countryCode"`
	UserAgent   string    `db:"user_agent"   json:"userAgent"`
}

type DBLike interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func dbMany(db DBLike, ctx context.Context, dest interface{}, query string, args ...interface{}) {
	err := db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		panic(err)
	}
}

func dbOne(db DBLike, ctx context.Context, dest interface{}, query string, args ...interface{}) bool {
	err := db.GetContext(ctx, dest, query, args...)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}

func dbExec(db DBLike, ctx context.Context, query string, arg interface{}) {
	_, err := db.NamedExecContext(ctx, query, arg)
	if err != nil {
		panic(err)
	}
}

func GetAccountByEmail(db DBLike, ctx context.Context, email string) (*Account, bool) {
	var account Account
	ok := dbOne(db, ctx, &account, `
		SELECT * FROM accounts WHERE email = $1
	`, email)
	return &account, ok
}

func CreateAccount(db DBLike, ctx context.Context, email, password string) *Account {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	account := Account{
		ID:             newID("acct_"),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Email:          email,
		HashedPassword: fmt.Sprintf("%s", hashed),
	}

	dbExec(db, ctx, `
		INSERT INTO accounts (id, created_at, updated_at, email, hashed_password) 
		VALUES (:id, :created_at, :updated_at, :email, :hashed_password) 
	`, &account)

	return &account
}

func FindAnalyticsEvents(db DBLike, ctx context.Context, websiteID string) []AnalyticsEvent {
	var events []AnalyticsEvent
	dbMany(db, ctx, &events, `
		SELECT * FROM analytics_events 
		WHERE website_id = $1
		ORDER BY created_at DESC
	`, websiteID)
	return events
}

func CreateAnalyticsEvent(db DBLike, ctx context.Context, websiteID, name, sessionKey string) *AnalyticsEvent {
	event := AnalyticsEvent{
		ID:         newID("evnt_"),
		WebsiteID:  websiteID,
		CreatedAt:  time.Now(),
		Name:       name,
		SessionKey: sessionKey,
	}

	dbExec(db, ctx, `
		INSERT INTO analytics_events (id, website_id, created_at, sid, name) 
		VALUES (:id, :website_id, :created_at, :sid, :name)
	`, &event)

	return &event
}

func FindAnalyticsPageviews(db DBLike, ctx context.Context, websiteID string) []AnalyticsPageview {
	var pageviews []AnalyticsPageview
	dbMany(db, ctx, &pageviews, `
		SELECT * FROM analytics_pageviews 
		WHERE website_id = $1
		ORDER BY created_at DESC
	`, websiteID)
	return pageviews
}

func CreateAnalyticsPageview(db DBLike, ctx context.Context, websiteID, host, path, screensize, country, sid, userAgent string) *AnalyticsPageview {
	pageview := AnalyticsPageview{
		ID:          newID("pgvw_"),
		CreatedAt:   time.Now(),
		WebsiteID:   websiteID,
		SID:         sid,
		Host:        host,
		Path:        path,
		ScreenSize:  screensize,
		CountryCode: country,
		UserAgent:   userAgent,
	}
	dbExec(db, ctx, `
		INSERT INTO analytics_pageviews (id, website_id, sid, created_at, host, path, screen_size, country_code, user_agent) 
		VALUES (:id, :website_id, :sid, :created_at, :host, :path, :screen_size, :country_code, :user_agent)
	`, &pageview)
	return &pageview
}

func CreateWebsite(db DBLike, ctx context.Context, accountID, domain string) *Website {
	website := Website{
		ID:        newID("site_"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		AccountId: accountID,
		Domain:    domain,
	}
	dbExec(db, ctx, `
		INSERT INTO websites (id, account_id, created_at, updated_at, domain) 
		VALUES (:id, :account_id, :created_at, :updated_at, :domain) 
	`, &website)
	return &website
}

func FindWebsitesByAccountID(db DBLike, ctx context.Context, accountID string) []Website {
	var websites []Website
	dbMany(db, ctx, &websites, `
		SELECT * FROM websites
		WHERE account_id = $1;
	`, accountID)
	return websites
}
