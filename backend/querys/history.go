package querys

const AddingHistoryRow = `
INSERT INTO HISTORYS (DATE, VALUE, ID_CREATOR, ID_PRODUCT, ID_AFILIATED, ID_TRANSACTION) 
VALUES (@` + NamedDate + `, @` + NamedValue + `, @` + NamedCreatorsId + `, @` + NamedProductId + `, @` + NamedAfiliatedId + `, @` + NamedTransactionId + `)
`
