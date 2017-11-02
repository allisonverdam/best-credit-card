package services

import (
	"time"

	"github.com/go-ozzo/ozzo-routing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/go-ozzo/ozzo-dbx"
)

func testDBCall(db *dbx.DB, f func(rs app.RequestScope, c routing.Context)) {
	rs := mockRequestScope(db)

	defer func() {
		rs.Tx().Rollback()
	}()

	// Mock up a routing context
	c := routing.NewContext(nil, nil)
	c.Set("Context", rs)

	f(rs, *c)
}

type requestScope struct {
	app.Logger
	tx        *dbx.Tx
	now       time.Time
	requestID string
	userID    int
}

func mockRequestScope(db *dbx.DB) app.RequestScope {
	tx, _ := db.Begin()
	return &requestScope{
		userID: 1,
		tx:     tx,
	}
}

func (rs *requestScope) UserID() int {
	return rs.userID
}

func (rs *requestScope) SetUserID(id int) {
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return "test"
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return false
}

func (rs *requestScope) SetRollback(v bool) {
}

func (rs *requestScope) Now() time.Time {
	return time.Now()
}
