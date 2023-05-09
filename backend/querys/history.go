package querys

const AddingHistoryRow = `
INSERT INTO HISTORIES (DATE, VALUE, CREATOR_ID, PRODUCT_ID, AFILIATED_ID, TRANSACTION_ID) 
VALUES (@` + NamedDate + `, @` + NamedValue + `, @` + NamedCreatorsId + `, @` + NamedProductId + `, @` + NamedAfiliatedId + `, @` + NamedTransactionId + `)
`
