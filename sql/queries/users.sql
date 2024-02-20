-- name: CreateUser :one
INSERT INTO Registeredusers (id, username, created_at, updated_at, email, password, whatsapp_number,apikey)
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
--
-- name: UserSigninout :exec
UPDATE Registeredusers 
SET loggedIn = $1
WHERE id = $2;
--
-- name: UpdateApiKey :exec
UPDATE Registeredusers
SET apikey = $1
WHERE id = $2;
--
-- name: GetUserDetails :one
SELECT * FROM Registeredusers
WHERE email = $1 OR whatsapp_number = $2;
--
-- name: GetUserByToken :one
SELECT * FROM Registeredusers
WHERE apikey = $1;
--
-- name: Updateusername :one
UPDATE RegisteredUsers
SET username = $1, updated_at = $2, updated_username = $3
WHERE id = $4
RETURNING *;

-- name: Updatepassword :one
UPDATE Registeredusers
SET password = $1,  updated_at = $2, updated_password = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUserDetail :exec
DELETE FROM RegisteredUsers
WHERE id = $1;
--

-- name: VerifyUserEmail :exec
UPDATE Registeredusers
SET isEmailVerified = $1, verified_at  =$2
WHERE id = $3;
--