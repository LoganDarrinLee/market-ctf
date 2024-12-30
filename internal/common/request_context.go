package common

import (
	"context"

	"github.com/rs/xid"
)

// Generate returns a new XID as a string.
func GenerateNewID() string {
	return xid.New().String()
}

// Request context will not contain user info.
type RequestContext struct {
	Ctx            context.Context
	RequestID      string
	SessionTokenID string
}

type contextKey string

const (
	RequestIDKey    contextKey = "requestID"
	SessionTokenKey contextKey = "sessionToken"
)

func NewRequestContext(ctx context.Context) *RequestContext {
	// Retrieve RequestID from context
	reqID, _ := ctx.Value(RequestIDKey).(string)
	tok, _ := ctx.Value(SessionTokenKey).(string)

	return &RequestContext{
		Ctx:            ctx,
		RequestID:      reqID,
		SessionTokenID: tok,
	}
}
