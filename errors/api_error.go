package errors

// APIError representa um erro que pode ser enviado na resposta da requisição
type APIError struct {
	// Status representa o status code da requisição
	Status int `json:"-" description:"Status Code do erro."`
	// ErrorCode é um codigo unico que identifica o erro
	ErrorCode string `json:"error_code" description:"Código do erro."`
	// Message é a mensagem de erro que pode ser exibida para o ususario
	Message string `json:"message" description:"Mensagem do erro."`
	// DeveloperMessage é a mensagem de erro destinada para os desenvolvedores
	DeveloperMessage string `json:"developer_message,omitempty" description:"Mensagem para os desenvolvedores."`
	// Details especifica as informações adicionais do erro
	Details interface{} `json:"details,omitempty" description:"Detalhes do erro."`
}

// Error retorna a mensagem do erro.
func (e APIError) Error() string {
	return e.Message
}

// StatusCode retorna o status code.
func (e APIError) StatusCode() int {
	return e.Status
}
