package querys

const AddAfiliate = `
INSERT INTO AFILIATEDS (NAME, LEFT_OVER ,CREATOR_ID)
SELECT @` + NamedName + `, 0.0, CREATORS.ID FROM CREATORS
WHERE CREATORS.USER_ID = @` + NamedUserId + ``

const GetAllAfiliates = `
SELECT AFILIATEDS.ID, AFILIATEDS.NAME FROM AFILIATEDS
INNER JOIN CREATORS ON CREATORS.ID = AFILIATEDS.CREATOR_ID
WHERE CREATORS.USER_ID = @` + NamedUserId + ``
