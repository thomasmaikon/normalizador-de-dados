package querys

const AddingHistoryRow = `
INSERT INTO HISTORIES (DATE, VALUE, CREATOR_ID, PRODUCT_ID, AFILIATED_ID, TRANSACTION_ID) 
VALUES (@` + NamedDate + `, @` + NamedValue + `, @` + NamedCreatorsId + `, @` + NamedProductId + `, @` + NamedAfiliatedId + `, @` + NamedTransactionId + `)
`

const GetAllDataFromUser = `
SELECT 
	AFILIATEDS.Name as afiliate,
	PRODUCTS.Description as product, 
	TRANSACTIONS.Description as transaction,
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

const GetAmmountReceivedValueAtCreator = `
SELECT COALESCE(SUM(VALUE),0) FROM HISTORIES h  
WHERE H.creator_id  = @` + NamedID + ` AND H.transaction_id = 1 or H.transaction_id = 2`

const GetAmmountPaidValueAtCreator = `
SELECT COALESCE(SUM(VALUE),0) FROM HISTORIES h  
WHERE H.creator_id  = @` + NamedID + ` AND H.transaction_id = 3`
