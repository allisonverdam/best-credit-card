package services

import (
	"time"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
	jwt "github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"golang.org/x/crypto/bcrypt"
)

// AuthService provides services related with auth.
type AuthService struct {
	dao personDAO
}

// NewAuthService creates a new AuthService with the given card DAO.
func NewAuthService(dao personDAO) *AuthService {
	return &AuthService{dao}
}

func (s *AuthService) Register(rs app.RequestScope, model *models.Person) (*models.Person, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

func (s *AuthService) Login(c *routing.Context, credential models.Credential, signingKey string) (string, error) {
	var daoPerson = daos.NewPersonDAO()

	identity := Authenticate(credential, daoPerson, c)
	if identity == nil {
		return "", errors.Unauthorized("invalid credential")
	}

	token, err := GenerateJWT(identity, signingKey)

	return token, err
}

func GenerateJWT(identity *models.Person, signingKey string) (string, error) {
	token, err := auth.NewJWT(jwt.MapClaims{
		"id":   identity.Id,
		"name": identity.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}, signingKey)
	if err != nil {
		return "", errors.Unauthorized(err.Error())
	}
	return token, nil
}

func Authenticate(c models.Credential, personDao *daos.PersonDAO, r *routing.Context) *models.Person {
	rs := app.GetRequestScope(r)
	person, err := personDao.GetPersonByUserName(rs, c.Username)
	if err != nil {
		return nil
	}

	if c.Username == person.Username && CheckPasswordHash(c.Password, person.Password) {
		return person
	}
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	userID := j.Claims.(jwt.MapClaims)["id"].(float64)
	app.GetRequestScope(c).SetUserID(int(userID))
	return nil
}
