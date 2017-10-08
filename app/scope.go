package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

//RequestScope contem a especificação das informações que passam nas requests da aplicação
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() int
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id int)
	// RequestID returns the ID of the current request
	RequestID() string
	//Tx retorna o database que está sendo usado na transação
	Tx() *dbx.Tx
	//SetTx indica o database que sera usado na transação
	SetTx(tx *dbx.Tx)
	//Rollback retorna um valor indicando se deve ser feito um rollback na transação atual
	Rollback() bool
	//SetRollback indica ser feito um rollback na transação atual
	SetRollback(bool)
	//Now retorna o timestamp representando o momento em que a requisição foi processada
	Now() time.Time
}

type requestScope struct {
	Logger              // looger com as informações da requisição
	now       time.Time // momento em que a requisição foi processada
	requestID string    // identifica a requisição HTTP
	userID    int       // identifica o usuario atual
	rollback  bool      // diz se é necessario reverter a requisição atual
	tx        *dbx.Tx   // transação atual
}

func (rs *requestScope) UserID() int {
	return rs.userID
}

func (rs *requestScope) SetUserID(id int) {
	rs.Logger.SetField("UserID", strconv.Itoa(id))
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope cria um novo RequestScope com as informações da requisição atual.
func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
