// Error handler for the web application
package common

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ErrorHandler interface {
	LogError(RequestContext, error) error
}

// Loop through slice of ErrorHandlers and log the error.
func LogError(rq RequestContext, e error, xe ...ErrorHandler) {
	for i := len(xe) - 1; i >= 0; i-- {
		if err := xe[i].LogError(rq, e); err != nil {
			log.Println(e)
		}
	}
}

// Database error logging
type DatabaseLogger struct {
	ConnectionPool *pgxpool.Conn
}

// Log error into database
func (d DatabaseLogger) LogError(rq RequestContext, e error) error {
	return nil
}

type TerminalLogger struct{}

// Log error to terminal
func (t TerminalLogger) LogError(rq RequestContext, e error) error {
	log.Println(e)
	return nil
}

// Client facing error handler.
type ClientLogger struct{}

// Log error to client, filtered to show only what they need to see.
func (c ClientLogger) LogError(rq RequestContext, e error) error {
	return nil
}
