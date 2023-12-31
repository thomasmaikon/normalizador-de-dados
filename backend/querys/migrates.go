package querys

const TransactionTypes = `
	INSERT INTO TRANSACTIONS (DESCRIPTION, KEY_FEATURE, SIGNAL)
	SELECT 'Venda produtor', 'Entrada', '+'
	WHERE NOT EXISTS (SELECT 1 FROM TRANSACTIONS WHERE ID = 1);

	INSERT INTO TRANSACTIONS (DESCRIPTION, KEY_FEATURE, SIGNAL)
	SELECT 'Venda afiliado', 'Entrada', '+'
	WHERE NOT EXISTS (SELECT 1 FROM TRANSACTIONS WHERE ID = 2);

	INSERT INTO TRANSACTIONS (DESCRIPTION, KEY_FEATURE, SIGNAL)
	SELECT 'Comissao paga', 'Saida', '-'
	WHERE NOT EXISTS (SELECT 1 FROM TRANSACTIONS WHERE ID = 3);

	INSERT INTO TRANSACTIONS (DESCRIPTION, KEY_FEATURE, SIGNAL)
	SELECT 'Comissao recebida', 'Entrada', '+'
	WHERE NOT EXISTS (SELECT 1 FROM TRANSACTIONS WHERE ID = 4);
`
