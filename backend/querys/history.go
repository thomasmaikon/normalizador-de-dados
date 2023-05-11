package querys

const AddingHistoryRow = `
INSERT INTO HISTORIES (DATE, VALUE, CREATOR_ID, PRODUCT_ID, AFILIATED_ID, TRANSACTION_ID) 
VALUES (@` + NamedDate + `, @` + NamedValue + `, @` + NamedCreatorsId + `, @` + NamedProductId + `, @` + NamedAfiliatedId + `, @` + NamedTransactionId + `)
`

const GetAllDataFromUser = `
SELECT 
	AFILIATEDS.Name as AfiliateName,
	PRODUCTS.Description as ProductDescription, 
	TRANSACTIONS.Description as TransactionDescription,
	HISTORIES.Value as Value,
	HISTORIES.Date as Date
FROM HISTORIES
INNER JOIN CREATORS 
	ON CREATORS.ID = HISTORIES.CREATOR_ID
INNER JOIN PRODUCTS 
	ON PRODUCTS.ID = HISTORIES.PRODUCT_ID
INNER JOIN TRANSACTIONS 
	ON TRANSACTIONS.ID = HISTORIES.TRANSACTION_ID
LEFT JOIN AFILIATEDS 
	ON AFILIATEDS.ID = HISTORIES.AFILIATED_ID
WHERE CREATORS.USER_ID = @` + NamedUserId + ``
