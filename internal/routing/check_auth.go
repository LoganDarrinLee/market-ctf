package routing

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
	"github.com/LoganDarrinLee/market-ctf/internal/db"
)

func checkAuth(q *db.Queries, rq *common.RequestContext, role string) (bool, db.GetUserWithSessionTokenRow, error) {
	// Grab user from session token
	userSessionRow, err := q.GetUserWithSessionToken(rq.Ctx, pgtype.Text{String: rq.SessionTokenID})
	if err != nil {
		return false, userSessionRow, err
	}

	// Get role
	roleRow, err := q.GetRoleWithName(rq.Ctx, role)

	if roleRow.ID == userSessionRow.RoleID.Int32 {
		// Access granted
		return true, userSessionRow, nil
	}

	// Default is to return false.
	return false, userSessionRow, err
}

func isAuthenticated() bool {
	return false
}
