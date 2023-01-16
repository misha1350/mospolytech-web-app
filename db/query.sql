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

