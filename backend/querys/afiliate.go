package querys

const AddAfiliate = `
INSERT INTO AFILIATEDS (NAME, LEFT_OVER ,CREATOR_ID)
SELECT @` + NamedName + `, 0.0, CREATORS.ID from LOGINS 
INNER JOIN USERS on USERS.login_id = logins.id 
INNER  JOIN CREATORS ON CREATORS.user_id  = users.id 
WHERE CREATORS.id = @` + NamedCreatorsId + ` AND logins.email = @` + NamedEmail + ``
