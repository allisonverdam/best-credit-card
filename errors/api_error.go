package errors

// APIError representa um erro que pode ser enviado na resposta da requisição
type APIError struct {
	// Status representa o status code da requisição
	Status int `json:"-"`
	// ErrorCode é um codigo unico que identifica o erro
	ErrorCode string `json:"error_code"`
	// Message é a mensagem de erro que pode ser exibida para o ususario
	Message string `json:"message"`
	// DeveloperMessage é a mensagem de erro destinada para os desenvolvedores
	DeveloperMessage string `json:"developer_message,omitempty"`
	// Details especifica as informações adicionais do erro
	Details interface{} `json:"details,omitempty"`
}

// Error retorna a mensagem do erro.
func (e APIError) Error() string {
	return e.Message
}

// StatusCode retorna o status code.
func (e APIError) StatusCode() int {
	return e.Status
}
