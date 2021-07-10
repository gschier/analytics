package main

import (
	"context"
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
	ID         string    `db:"id"          json:"id"`
	CreatedAt  time.Time `db:"created_at"  json:"createdAt"`
	Host       string    `db:"host"        json:"host"`
	Path       string    `db:"path"        json:"path"`
	UserAgent  string    `db:"user_agent"  json:"userAgent"`
	ScreenSize string    `db:"screen_size" json:"screenSize"`
	TimeZone   string    `db:"time_zone"   json:"timeZone"`
}

// type AnalyticsEventBucket struct {
// 	ID        string    `db:"id"          json:"id"`
// 	WebsiteID string    `db:"website_id"  json:"websiteId"`
// 	CreatedAt time.Time `db:"created_at"  json:"createdAt"`
// 	Start     time.Time `db:"start"       json:"start"`
// 	End       time.Time `db:"end"         json:"end"`
// 	Name      string    `db:"name"        json:"name"`
// }
//
// type AnalyticsPageviewBucket struct {
// 	ID        string    `db:"id"          json:"id"`
// 	WebsiteID string    `db:"website_id"  json:"websiteId"`
// 	CreatedAt time.Time `db:"created_at"  json:"createdAt"`
// 	Start     time.Time `db:"start"       json:"start"`
// 	End       time.Time `db:"end"         json:"end"`
// 	Host      string    `db:"host"        json:"host"`
// 	Path      string    `db:"path"        json:"path"`
// }

// type HLLColumn struct {
// 	Sketch *hyperloglog.Sketch
// }
//
// func (h HLLColumn) Value() (driver.Value, error) {
// 	return h.Sketch.MarshalBinary()
// }
//
// func (h HLLColumn) Scan(v interface{}) error {
// 	sketch := hyperloglog.New()
// 	err := sketch.UnmarshalBinary(v.([]byte))
// 	if err != nil {
// 		return err
// 	}
// 	h.Sketch = sketch
// 	return nil
// }

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

func (s *dbStore) ListAnalyticsEvents(ctx context.Context) ([]AnalyticsEvent, error) {
	var events []AnalyticsEvent
	err := s.db.SelectContext(ctx, &events, `SELECT * FROM analytics_events ORDER BY created_at DESC`)
	return events, err
}

func (s *dbStore) CreateAnalyticsEvent(ctx context.Context, websiteID, name string) (*AnalyticsEvent, error) {
	var event AnalyticsEvent
	err := s.db.QueryRowxContext(ctx, `
		INSERT INTO analytics_events (website_id, name) VALUES ($1, $2)
		RETURNING *
	`, websiteID, name).StructScan(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (s *dbStore) FindAnalyticsPageviews(ctx context.Context) ([]AnalyticsPageview, error) {
	var pageviews []AnalyticsPageview
	err := s.db.SelectContext(ctx, &pageviews, `SELECT * FROM analytics_pageviews ORDER BY created_at DESC`)
	return pageviews, err
}

func (s *dbStore) CreateAnalyticsPageview(ctx context.Context) (*AnalyticsPageview, error) {
	var pageview AnalyticsPageview
	err := s.db.QueryRowxContext(ctx, `
		INSERT INTO analytics_pageviews DEFAULT VALUES
		RETURNING *
	`).StructScan(&pageview)
	if err != nil {
		return nil, err
	}

	return &pageview, nil
}

func (s *dbStore) CreateWebsite(ctx context.Context, accountID, name string) (*Website, error) {
	var website Website
	err := s.db.QueryRowxContext(ctx, `
		INSERT INTO websites (account_id, domain) VALUES ($1, $2) 
		RETURNING *
	`, accountID, name).StructScan(&website)
	if err != nil {
		return nil, err
	}

	return &website, nil
}

func (s *dbStore) FindWebsitesByAccountID(ctx context.Context, accountID string) ([]Website, error) {
	var websites []Website
	err := s.db.SelectContext(ctx, &websites, `
		SELECT * FROM websites
		WHERE account_id = $1;
	`, accountID)
	if err != nil {
		return nil, err
	}

	return websites, nil
}
