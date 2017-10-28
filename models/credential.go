package models

type Credential struct {
	Username string `json:"username" description:"Usuario."`
	Password string `json:"password" description:"Senha."`
}
