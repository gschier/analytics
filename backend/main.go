package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println()
	fmt.Printf("\u001B[32;1m*******************************\u001B[0m\n")
	fmt.Printf("\u001B[32;1m*          analytics          *\u001B[0m\n")
	fmt.Printf("\u001B[32;1m*******************************\u001B[0m\n")

	InitConfig()

	if Config.MigrateOnStart {
		mustMigrate(context.Background(), GetDB())
	}

	websiteID := ensureDummyWebsite()
	fmt.Println("[main] Website", websiteID)
	logger := NewLogger("router")

	h := SetupRouter()
	loggedRouter := LoggingMiddleware(logger)(h)
	fmt.Printf("[schier.co] \033[32;1mStarted server on http://%s:%s\033[0m\n", Config.Host, Config.Port)
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, loggedRouter))
}

func ensureDummyWebsite() string {
	account, accountExists := GetAccountByEmail(GetDB(), context.Background(), "greg@schier.co")

	if accountExists {
		websites := FindWebsitesByAccountID(GetDB(), context.Background(), account.ID)
		return websites[0].ID
	}

	a := CreateAccount(GetDB(), context.Background(), "greg@schier.co", "my-pass!")
	w := CreateWebsite(GetDB(), context.Background(), a.ID, "schier.co")

	return w.ID
}

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: http.StatusOK}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(
						fmt.Sprintf("%s\n%s", err, debug.Stack()),
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Debug(
				"Request completed to "+r.URL.EscapedPath(),
				"status", wrapped.status,
				"headers", r.Header,
				"addr", r.RemoteAddr,
				"query", r.URL.Query(),
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}
