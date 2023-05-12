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

	ErrorCodeFaildCreateCreator    = 7
	ErrorMessageFaildCreateCreator = "Falha ao cadastrar um creator"

	ErrorCodeFaildCreatorNotFound    = 8
	ErrorMessageFaildCreatorNotFound = "Falha ao buscar creator"

	ErrorCodeFaildAddTransaction    = 9
	ErrorMessageFaildAddTransaction = "Falha ao adicionar transacao"

	ErrorCodeFaildNormalizeFile    = 10
	ErrorMessageFaildNormalizeFile = "Falha ao fazer normalizacao dos dados do arquivo"

	ErrorCodeFindProduct    = 11
	ErrorMessageFindProduct = "Falha ao procurar produto"

	ErrorCodeCreateProduct    = 12
	ErrorMessageCreateProduct = "Falha ao criar produto"

	ErrorCodeFaildNotFoundALlProducts    = 13
	ErrorMessageFaildNotFoundALlProducts = "Falha ao buscar todos os produtos"

	ErrorCodeFaildCreateLogin    = 14
	ErrorMessageFaildCreateLogin = "Email ja cadastrado"

	ErrorCodeFaildLogin    = 15
	ErrorMessageFaildLogin = "Credenciais invalidas"

	ErrorCodeHistorical    = 16
	ErrorMessageHistorical = "Nao foi possivel encontrar historico de transacoes"

	ErrorCodeAmountFromCreator    = 17
	ErrorMessageAmountFromCreator = "Falha ao calcular valor total do creator"

	ErrorCodeAmountFromAfiliate    = 18
	ErrorMessageAmountFromAfiliate = "Falha ao calcular valor total do afiliado"

	ErrorCodeNotFoundHistoricalFromAfiliate    = 19
	ErrorMessageNotFoundHistoricalFromAfiliate = "Falha ao buscar historico do afiliado"

	ErrorCodeCouldNotCreateUser    = 20
	ErrorMessageCouldNotCreateUser = "Nao foi possivel cadastrar o usuario"

	ErrorCodeNotFidUser    = 21
	ErrorMessageNotFidUser = "Nao foi possivel encontrar o usuario"
)
