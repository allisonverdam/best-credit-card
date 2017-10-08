package app

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/fault"
)

// Transactional retorna um handler que junta os handlers aninhados com uma transação de banco de dados.
// Se um handler aninhado retornar um erro ou um "panic" acontecer, ele irá reverter a transação.
// Caso contrário, irá confirmar a transação depois que os handlers aninhados terminem a execução.
// Ao executar app.Context.SetRollback(true), será solicitado a reversão da transação.
func Transactional(db *dbx.DB) routing.Handler {
	return func(c *routing.Context) error {
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		rs := GetRequestScope(c)
		rs.SetTx(tx)

		err = fault.PanicHandler(rs.Errorf)(c)

		var e error
		if err != nil || rs.Rollback() {
			// reverte a transação se tiver um erro ou se for requisitado um rollback
			e = tx.Rollback()
		} else {
			e = tx.Commit()
		}

		if e != nil {
			if err == nil {
				// o erro vai ser registrado pelo nosso logger
				return e
			}
			// registra apensa o erro da transação
			rs.Error(e)
		}

		return err
	}
}
