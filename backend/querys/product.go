package querys

const CreateNewProduct = `
INSERT INTO PRODUCTS (description, price, creator_id) 
SELECT @` + NamedDescription + `, @` + NamedPrice + `, CREATORS.ID from LOGINS 
INNER JOIN USERS on USERS.login_id = logins.id 
INNER  JOIN CREATORS ON CREATORS.user_id  = users.id 
WHERE CREATORS.id = @` + NamedCreatorsId + ` AND logins.email = @` + NamedEmail + ``
