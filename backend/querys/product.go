package querys

const CreateNewProduct = `
INSERT INTO PRODUCTS (description, price, creator_id) 
SELECT @` + NamedDescription + `, @` + NamedPrice + `, CREATORS.ID FROM USERS
INNER  JOIN CREATORS ON CREATORS.USER_ID  = USERS.ID 
WHERE CREATORS.ID = @` + NamedCreatorsId + ` AND CREATORS.USER_ID = @` + NamedUserId + ``
