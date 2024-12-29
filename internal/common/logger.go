package common

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Logger interface {
	// Log message to terminal
	WriteInfo(rq *RequestContext, msg string)

	// Log error to terminal.
	WriteError(rq *RequestContext, err error)

	// Save error to database.
	SaveError(rq *RequestContext, err error, conn *pgxpool.Conn) error
}

type BasicLogger struct{}

// Format log message with RequestContext information
func (b *BasicLogger) formatMessage(rc *RequestContext, level, msg string) string {
	return fmt.Sprintf(
		"[%s] Request ID: %s | %s",
		level, rc.RequestID, msg,
	)
}

func (b *BasicLogger) WriteInfo(rc *RequestContext, msg string) {
	log.Println(b.formatMessage(rc, "INFO", msg))
}

func (b *BasicLogger) WriteError(rc *RequestContext, msg string, err error) {
	log.Println(b.formatMessage(rc, "ERROR", fmt.Sprintf("%s - %v", msg, err)))
}

func (b *BasicLogger) SaveError(rq *RequestContext, err error, conn *pgxpool.Conn) error { return nil }
