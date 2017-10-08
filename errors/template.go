package errors

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type (
	// Params is used to replace placeholders in an error template with the corresponding values.
	Params map[string]interface{}

	errorTemplate struct {
		Message          string `yaml:"message"`
		DeveloperMessage string `yaml:"developer_message"`
	}
)

var templates map[string]errorTemplate

// LoadMessages carrega o arquivo que contem os templates de mensagens de erro
func LoadMessages(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	templates = map[string]errorTemplate{}
	return yaml.Unmarshal(bytes, &templates)
}

// NewAPIError cria um erro customizado com o status code, error code, e os parâmetros e substitui com os valores do template de erros.
// Se o parâmetro for nulo fica com o valor padrão do erro.
func NewAPIError(status int, code string, params Params) *APIError {
	err := &APIError{
		Status:    status,
		ErrorCode: code,
		Message:   code,
	}

	if template, ok := templates[code]; ok {
		err.Message = template.getMessage(params)
		err.DeveloperMessage = template.getDeveloperMessage(params)
	}

	return err
}

// getMessage retorna a mensagem de erro substituindo os espaços reservados no template de erros com os parâmetros.
func (e errorTemplate) getMessage(params Params) string {
	return replacePlaceholders(e.Message, params)
}

// getDeveloperMessage retorna a mensagem do desenvolvedor substituindo os espaços reservados no template de erros com os parâmetros.
func (e errorTemplate) getDeveloperMessage(params Params) string {
	return replacePlaceholders(e.DeveloperMessage, params)
}

func replacePlaceholders(message string, params Params) string {
	if len(message) == 0 {
		return ""
	}
	for key, value := range params {
		message = strings.Replace(message, "{"+key+"}", fmt.Sprint(value), -1)
	}
	return message
}
