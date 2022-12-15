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