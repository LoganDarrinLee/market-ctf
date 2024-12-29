package common

import (
	"context"

	"github.com/rs/xid"
)

// Generate returns a new XID as a string.
func GenerateNewID() string {
	return xid.New().String()
}

type RequestContext struct {
	Ctx       context.Context
	RequestID string
}

type contextKey string

const (
	RequestIDKey contextKey = "requestID"
)

func NewRequestContext(ctx context.Context) *RequestContext {
	// Retrieve RequestID from context
	reqID, _ := ctx.Value(RequestIDKey).(string)

	return &RequestContext{Ctx: ctx, RequestID: reqID}
}
