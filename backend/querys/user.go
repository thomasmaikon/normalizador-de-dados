package querys

const FinUserByLoginId = `
SELECT * FROM USERS 
WHERE  LOGIN_ID = @` + NamedID + ``
