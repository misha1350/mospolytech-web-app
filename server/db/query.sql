-- name: GetCountry :one
SELECT * FROM countries
WHERE ID = ? LIMIT 1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY Name;

-- name: CreateCountry :execresult
INSERT INTO countries (
  Name
) VALUES (
  ?
);

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE Id = ?;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY ID;

-- name: GetUserToUpdate :one
SELECT * FROM users
WHERE ID = ? LIMIT 1;

-- name: GetOffices :many
SELECT * FROM offices
ORDER BY ID;

-- name: AddUser :execresult
INSERT INTO users (
  RoleID, Email, Password, FirstName, LastName, OfficeID, Birthdate, Active
) VALUES (
  1, ?, ?, ?, ?, ?, ?, 1
);

-- name: UpdateUser :exec
UPDATE users
SET
  RoleID = ?,
  Email = ?,
  FirstName = ?,
  LastName = ?,
  OfficeID = ?
WHERE ID = ?;

-- name: BanUser :exec
UPDATE users
SET Active = 0
WHERE ID = ?;

-- name: UnbanUser :exec
UPDATE users
SET Active = 1
WHERE ID = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE ID = ?;

-- -- name: FindUserByEmail :one
-- SELECT * FROM users
-- WHERE Email = ? LIMIT 1;

-- name: GetTracking :many
SELECT * FROM tracking
ORDER BY ID;

-- name: AddTracking :execresult
INSERT INTO tracking (
  UserID, Date, TimeIn, TimeOut, Hours, Notes
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: UpdateTracking :exec
UPDATE tracking
SET
  Date = ?,
  TimeIn = ?,
  TimeOut = ?,
  Hours = ?,
  Notes = ?
WHERE ID = ?;

-- name: DeleteTracking :exec
DELETE FROM tracking
WHERE ID = ?;

-- name: GetUserData :one
SELECT ID, RoleID, Email, Password, OfficeID, FirstName, LastName, Active FROM users
WHERE Email = ? LIMIT 1;

-- name: GetUserDataByID :one
SELECT ID, RoleID, Email, Password, OfficeID, FirstName, LastName, Active FROM users
WHERE ID = ? LIMIT 1;

-- name: GetAuthToken :one
SELECT TokenID, UserID, AuthToken FROM authentication_tokens
WHERE UserID = ? ORDER BY GeneratedAt DESC LIMIT 1;

-- name: InsertAuthToken :exec
INSERT INTO authentication_tokens (
  UserID, AuthToken, GeneratedAt, ExpiresAt
) VALUES (
  ?, ?, ?, ?
);