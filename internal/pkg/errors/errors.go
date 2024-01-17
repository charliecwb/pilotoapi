package errors

type ErrorResponse struct {
	Erro    string `json:"error"`
	Message string `json:"message"`
}

const (
	ErrFalhaConsultarDados = "Falha ao consultar dados"
	ErrParameterInvalido   = "Parâmetros inválidos"
	ErrRequestBodyInvalido = "Request body inválido"
	ErrFalhaEnviarDados    = "Falha ao enviar dados"
)
