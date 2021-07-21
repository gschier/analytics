package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
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

type Website struct {
	ID        string    `db:"id"         json:"id"`
	AccountId string    `db:"account_id" json:"accountId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Domain    string    `db:"domain"     json:"domain"`
}

type AnalyticsEvent struct {
	ID        string    `db:"id"          json:"id"`
	WebsiteID string    `db:"website_id"  json:"websiteId"`
	CreatedAt time.Time `db:"created_at"  json:"createdAt"`
	SID       string    `db:"sid"         json:"sid"`
	Name      string    `db:"name"        json:"name"`
}

type AnalyticsPageview struct {
	ID          string    `db:"id"           json:"id"`
	SID         string    `db:"sid"          json:"sid"`
	WebsiteID   string    `db:"website_id"   json:"websiteId"`
	CreatedAt   time.Time `db:"created_at"   json:"createdAt"`
	Host        string    `db:"host"         json:"host"`
	Path        string    `db:"path"         json:"path"`
	ScreenSize  string    `db:"screen_size"  json:"screenSize"`
	CountryCode string    `db:"country_code" json:"countryCode"`
	UserAgent   string    `db:"user_agent"   json:"userAgent"`
}

func GetAccountByEmail(db sqlx.QueryerContext, ctx context.Context, email string) (*Account, bool) {
	var account Account
	err := sqlx.GetContext(ctx, db, &account, `
		SELECT * FROM accounts WHERE email = $1
	`, email)
	if err == sql.ErrNoRows {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	return &account, true
}

func CreateAccount(db sqlx.ExtContext, ctx context.Context, email, password string) *Account {
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

	_, err = sqlx.NamedExecContext(ctx, db, `
		INSERT INTO accounts (id, created_at, updated_at, email, hashed_password)
		VALUES (:id, :created_at, :updated_at, :email, :hashed_password)
	`, &account)
	if err != nil {
		panic(err)
	}

	return &account
}

func CreateAnalyticsEvent(db sqlx.ExtContext, ctx context.Context, id, sid, websiteID, name string) *AnalyticsEvent {
	event := AnalyticsEvent{
		ID:        id,
		SID:       sid,
		WebsiteID: websiteID,
		CreatedAt: time.Now(),
		Name:      name,
	}

	_, err := sqlx.NamedExecContext(ctx, db, `
		INSERT INTO analytics_events (id, sid, website_id, created_at, name) 
		VALUES (:id, :sid, :website_id, :created_at, :name)
	`, &event)
	if err != nil {
		panic(err)
	}

	return &event
}

func FindAnalyticsPageviews(db sqlx.QueryerContext, ctx context.Context, websiteID string) []AnalyticsPageview {
	var pageviews []AnalyticsPageview
	err := sqlx.SelectContext(ctx, db, &pageviews, `
		SELECT * FROM analytics_pageviews 
		WHERE website_id = $1
		ORDER BY created_at DESC
		LIMIT 20
	`, websiteID)
	if err != nil {
		panic(err)
	}
	return pageviews
}

func FindAnalyticsPageviewsBuckets(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, websiteID string) []Bucket {
	type count struct {
		Total  int64     `db:"count_total"`
		Unique int64     `db:"count_unique"`
		Bucket time.Time `db:"bucket"`
	}

	var counts []count
	err := sqlx.SelectContext(ctx, db, &counts, `
		SELECT COUNT(id)                                                                                AS count_total,
			   COUNT(DISTINCT sid)                                                                      AS count_unique,
			   TO_TIMESTAMP(FLOOR((EXTRACT('epoch' FROM created_at) / 3600)) * 3600) AT TIME ZONE 'UTC' AS bucket
		FROM analytics_pageviews
		WHERE website_id = $1 AND created_at >= $2 AND created_at < $3
		GROUP BY bucket;
	`, websiteID, start, end)
	if err != nil {
		panic(err)
	}

	// Iterate over all buckets, to make sure we create ones for periods with no events
	buckets := make([]Bucket, end.Sub(start)/time.Hour)
	startFloored := GetBucketStart(start, PeriodHour)
	for i := 0; i < len(buckets); i++ {
		bucketStart := startFloored.Add(time.Duration(i) * time.Hour)
		bucket := Bucket{Start: bucketStart, End: bucketStart.Add(time.Hour)}

		// Try to find a count for the bucket. If so, fill it in
		for _, c := range counts {
			if !c.Bucket.Equal(bucket.Start) {
				continue
			}

			bucket.Total = c.Total
			bucket.Unique = c.Unique
		}

		buckets[i] = bucket
	}

	return buckets
}

type PathCount struct {
	Total  int64  `db:"count_total" json:"total"`
	Unique int64  `db:"count_unique" json:"unique"`
	Path   string `db:"path" json:"path"`
	Host   string `db:"host" json:"host"`
}

type ThingsCount struct {
	Total      int64   `db:"count_total" json:"total"`
	Unique     int64   `db:"count_unique" json:"unique"`
	Host       *string `db:"host" json:"host"`
	Path       *string `db:"path" json:"path"`
	Country    *string `db:"country_code" json:"country"`
	ScreenSize *string `db:"screen_size" json:"screenSize"`
}

func FindAnalyticsPageviewsPopularPages(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, websiteID string) []PathCount {
	var counts []PathCount

	err := sqlx.SelectContext(ctx, db, &counts, `
		SELECT path,
		       host,
			   COUNT(id)           AS count_total,
			   COUNT(DISTINCT sid) AS count_unique
		FROM analytics_pageviews
		WHERE website_id = $1
		  AND created_at >= $2
		  AND created_at < $3
		GROUP BY path, host
		ORDER BY count_unique DESC
		LIMIT 10;
	`, websiteID, start, end)
	if err != nil {
		panic(err)
	}

	return counts
}

func FindAnalyticsPageviewsPopularThings(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, websiteID string) []ThingsCount {
	var counts []ThingsCount

	err := sqlx.SelectContext(ctx, db, &counts, `
		SELECT screen_size, 
		       path,
		       host,
		       country_code,
			   COUNT(id)           AS count_total,
			   COUNT(DISTINCT sid) AS count_unique
		FROM analytics_pageviews
		WHERE country_code != '' 
		  AND website_id = $1
		  AND created_at >= $2
		  AND created_at < $3
		GROUP BY GROUPING SETS (screen_size, country_code, (path, host))
		ORDER BY count_unique DESC
		LIMIT 50;
	`, websiteID, start, end)
	if err != nil {
		panic(err)
	}

	return counts
}

func CreateAnalyticsPageview(db sqlx.ExtContext, ctx context.Context, id, sid, websiteID, host, path, screensize, country, userAgent string) *AnalyticsPageview {
	pageview := AnalyticsPageview{
		ID:          id,
		CreatedAt:   time.Now(),
		WebsiteID:   websiteID,
		SID:         sid,
		Host:        host,
		Path:        path,
		ScreenSize:  screensize,
		CountryCode: country,
		UserAgent:   userAgent,
	}
	_, err := sqlx.NamedExecContext(ctx, db, `
		INSERT INTO analytics_pageviews (id, website_id, sid, created_at, host, path, screen_size, country_code, user_agent) 
		VALUES (:id, :website_id, :sid, :created_at, :host, :path, :screen_size, :country_code, :user_agent)
	`, &pageview)
	if err != nil {
		panic(err)
	}
	return &pageview
}

func CreateWebsite(db sqlx.ExtContext, ctx context.Context, accountID, domain string) *Website {
	website := Website{
		ID:        newID("site_"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		AccountId: accountID,
		Domain:    domain,
	}
	_, err := sqlx.NamedExecContext(ctx, db, `
		INSERT INTO websites (id, account_id, created_at, updated_at, domain) 
		VALUES (:id, :account_id, :created_at, :updated_at, :domain) 
	`, &website)
	if err != nil {
		panic(err)
	}
	return &website
}

func FindWebsitesByAccountID(db sqlx.QueryerContext, ctx context.Context, accountID string) []Website {
	var websites []Website
	err := sqlx.SelectContext(ctx, db, &websites, `
		SELECT * FROM websites
		WHERE account_id = $1;
	`, accountID)
	if err != nil {
		panic(err)
	}
	return websites
}
