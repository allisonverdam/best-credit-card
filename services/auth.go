package services

import (
	"net/http"
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

func (s *AuthService) Register(rs app.RequestScope, person *models.Person) (*models.Person, error) {
	var err error
	if err = person.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.CreatePerson(rs, person); err != nil {
		return nil, err
	}

	personDB, err := s.dao.GetWithoutPassword(rs, person.Id)
	if err != nil {
		return nil, err
	}

	tempWallet := &models.Wallet{}
	tempWallet.PersonId = personDB.Id

	_, err = NewWalletService(daos.NewWalletDAO()).CreateWallet(rs, tempWallet)
	if err != nil {
		return nil, err
	}

	return personDB, err
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

func VerifyPersonOwner(rs app.RequestScope, id int, resource string) error {
	if rs.UserID() != id {
		return errors.NewAPIError(http.StatusForbidden, "FORBIDDEN", errors.Params{"message": "You're not allowed to do this.", "developer_message": "This " + resource + " does not belong to the authenticated user."})
	}
	return nil
}
