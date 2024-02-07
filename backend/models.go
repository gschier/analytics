package main

import (
	"context"
	"database/sql"
	"errors"
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
	ID          string    `db:"id"           json:"id"`
	WebsiteID   string    `db:"website_id"   json:"websiteId"`
	CreatedAt   time.Time `db:"created_at"   json:"createdAt"`
	SID         string    `db:"sid"          json:"sid"`
	Name        string    `db:"name"         json:"name"`
	Attributes  string    `db:"attributes"   json:"attributes"`
	ScreenSize  string    `db:"screen_size"  json:"screenSize"`
	CountryCode string    `db:"country_code" json:"countryCode"`
	Platform    string    `db:"platform"     json:"platform"`
	Version     string    `db:"version"      json:"version"`
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
	Referrer    string    `db:"referrer"     json:"referrer"`
}

func GetAccountByEmail(db sqlx.QueryerContext, ctx context.Context, email string) (*Account, bool) {
	var account Account
	err := sqlx.GetContext(ctx, db, &account, `
		SELECT * FROM accounts WHERE email = $1
	`, email)
	if errors.Is(err, sql.ErrNoRows) {
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

func CreateAnalyticsEvent(db sqlx.ExtContext, ctx context.Context, event *AnalyticsEvent) *AnalyticsEvent {
	event.CreatedAt = time.Now()

	_, err := sqlx.NamedExecContext(ctx, db, `
		INSERT INTO analytics_events (id, sid, website_id, created_at, name, attributes, screen_size, country_code, platform, version) 
		VALUES (:id, :sid, :website_id, :created_at, :name, :attributes, :screen_size, :country_code, :platform, :version)
	`, &event)
	if err != nil {
		panic(err)
	}

	return event
}

//func FindAnalyticsPageviews(db sqlx.QueryerContext, ctx context.Context, websiteID string) []AnalyticsPageview {
//	var pageviews []AnalyticsPageview
//	err := sqlx.SelectContext(ctx, db, &pageviews, `
//		SELECT * FROM analytics_pageviews
//		WHERE website_id = $1
//		ORDER BY created_at DESC
//		LIMIT 20
//	`, websiteID)
//	if err != nil {
//		panic(err)
//	}
//	return pageviews
//}

func CountAnalyticsPageviewsRecent(db sqlx.QueryerContext, ctx context.Context, websiteID string) int64 {
	var count int64
	err := sqlx.GetContext(ctx, db, &count, `
		SELECT COUNT(DISTINCT sid) FROM analytics_pageviews 
		WHERE website_id = $1 AND created_at > (NOW() - INTERVAL '5 minutes')
		LIMIT 20
	`, websiteID)
	if err != nil {
		panic(err)
	}
	return count
}

func FindAnalyticsPageviewsBuckets(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, period time.Duration, websiteID string) []Bucket {
	var counts []struct {
		Total  int64     `db:"count_total"`
		Unique int64     `db:"count_unique"`
		Bucket time.Time `db:"bucket"`
	}

	startCeil := CeilToPeriod(start, period)
	endCeil := CeilToPeriod(end, period)

	err := sqlx.SelectContext(ctx, db, &counts, `
		SELECT COUNT(id)                                                                            AS count_total,
			   COUNT(DISTINCT sid)                                                                  AS count_unique,
			   TO_TIMESTAMP(FLOOR((EXTRACT('epoch' FROM created_at) / $4)) * $4) AT TIME ZONE 'UTC' AS bucket
		FROM analytics_pageviews
		WHERE website_id = $1 AND created_at >= $2 AND created_at < $3
		GROUP BY bucket;
	`, websiteID, startCeil, endCeil, int(period.Seconds()))
	if err != nil {
		panic(err)
	}

	// Iterate over all buckets, to make sure we create ones for periods with no events
	buckets := make([]Bucket, end.Sub(start)/period)

	for i := 0; i < len(buckets); i++ {
		bucketStart := startCeil.Add(time.Duration(i) * period)
		bucket := Bucket{
			Start: bucketStart,
			End:   bucketStart.Add(period),
		}

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

type PopularEventCount struct {
	Total      int64   `db:"count_total" json:"total"`
	Unique     int64   `db:"count_unique" json:"unique"`
	Name       *string `db:"name" json:"name"`
	Platform   *string `db:"platform" json:"platform"`
	Version    *string `db:"version" json:"version"`
	Country    *string `db:"country_code" json:"country"`
	ScreenSize *string `db:"screen_size" json:"screenSize"`
}

func FindAnalyticsEventsPopular(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, websiteID string) []PopularEventCount {
	var counts []PopularEventCount
	err := sqlx.SelectContext(ctx, db, &counts, `
		SELECT  name,
			   COUNT(id)           AS count_total,
			   COUNT(DISTINCT sid) AS count_unique
		FROM analytics_events
		WHERE 
		  website_id = $1
		  AND created_at >= $2
		  AND created_at < $3
		GROUP BY name
		ORDER BY count_unique DESC
		LIMIT 50;
	`, websiteID, start, end)
	if err != nil {
		panic(err)
	}

	return counts
}

type PopularCount struct {
	Total      int64   `db:"count_total" json:"total"`
	Unique     int64   `db:"count_unique" json:"unique"`
	Host       *string `db:"host" json:"host"`
	Path       *string `db:"path" json:"path"`
	Country    *string `db:"country_code" json:"country"`
	ScreenSize *string `db:"screen_size" json:"screenSize"`
}

func FindAnalyticsPageviewsPopular(db sqlx.QueryerContext, ctx context.Context, start, end time.Time, websiteID string) []PopularCount {
	var counts []PopularCount
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
		GROUP BY GROUPING SETS (
		  (screen_size), 
		  (country_code), 
		  (path, host), 
		  ()
	  	)
		ORDER BY count_unique DESC
		LIMIT 500;
	`, websiteID, start, end)
	if err != nil {
		panic(err)
	}

	return counts
}

func CreateAnalyticsPageview(db sqlx.ExtContext, ctx context.Context, pageview *AnalyticsPageview) *AnalyticsPageview {
	pageview.CreatedAt = time.Now()
	_, err := sqlx.NamedExecContext(ctx, db, `
		INSERT INTO analytics_pageviews (id, website_id, sid, created_at, host, path, screen_size, country_code, user_agent, referrer) 
		VALUES (:id, :website_id, :sid, :created_at, :host, :path, :screen_size, :country_code, :user_agent, :referrer)
	`, pageview)
	if err != nil {
		panic(err)
	}
	return pageview
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
