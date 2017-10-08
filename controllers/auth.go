package controllers

import (
	"time"

	"github.com/allisonverdam/go-api-mcc/app"
	"github.com/allisonverdam/go-api-mcc/daos"
	"github.com/allisonverdam/go-api-mcc/errors"
	"github.com/allisonverdam/go-api-mcc/models"
	jwt "github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		var daoPerson = daos.NewPersonDAO()
		var credential Credential
		if err := c.Read(&credential); err != nil {
			return errors.Unauthorized(err.Error())
		}

		identity := authenticate(credential, daoPerson, c)
		if identity == nil {
			return errors.Unauthorized("invalid credential")
		}

		token, err := auth.NewJWT(jwt.MapClaims{
			"id":   identity.Id,
			"name": identity.Name,
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		}, signingKey)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		return c.Write(map[string]string{
			"token": token,
		})
	}
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	userID := j.Claims.(jwt.MapClaims)["id"].(float64)
	app.GetRequestScope(c).SetUserID(int(userID))
	return nil
}

func authenticate(c Credential, personDao *daos.PersonDAO, r *routing.Context) *models.Person {
	rs := app.GetRequestScope(r)
	person, err := personDao.GetByPersonName(rs, c.Username)
	if err != nil {
		return nil
	}

	if c.Username == person.Username && c.Password == person.Password {
		return person
	}
	return nil
}
