package querys

const AddAfiliate = `
INSERT INTO AFILIATEDS (NAME, LEFT_OVER ,CREATOR_ID)
SELECT @` + NamedName + `, 0.0, CREATORS.ID FROM CREATORS
WHERE CREATORS.USER_ID = @` + NamedUserId + ``
