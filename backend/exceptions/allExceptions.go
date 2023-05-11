package exceptions

const (
	ErrorCodeCreateUser    = 1
	ErrorMessageCreateUser = "Falha ao cadastrar usuario"

	ErrorCodeConvertToUserId   = 2
	ErrorMessageConverToUserId = "Falha ao identificar numero do usuario"

	ErrorCodeConvertToAfiliateId  = 3
	ErrorMessageConvertAfiliateId = "Falha ao receber afiliado"

	ErrorCodeAfiliateNotFound    = 4
	ErrorMessageAfiliateNotFound = "Afiliado nao encontrado"

	ErrorCodeAddAfiliate   = 5
	ErrorMessageddAfiliate = "Falha ao adicionar afiliado"

	ErrorCodeAllAfiliateNotFound    = 6
	ErrorMessageAllAfiliateNotFound = "Falha ao buscar todos afiliados"
)
