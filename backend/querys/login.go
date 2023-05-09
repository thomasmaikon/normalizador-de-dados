package querys

const CreateLogin = `
INSERT INTO LOGINS (EMAIL, PASSWORD) 
VALUES (@` + NamedEmail + `, @` + NamedPassword + `)
`
