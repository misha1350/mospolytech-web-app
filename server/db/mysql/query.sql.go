// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package mysql

import (
	"context"
	"database/sql"
	"time"
)

const addTracking = `-- name: AddTracking :execresult
INSERT INTO tracking (
  UserID, Date, TimeIn, TimeOut, Hours, Notes
) VALUES (
  ?, ?, ?, ?, ?, ?
)
`

type AddTrackingParams struct {
	Userid  int32
	Date    time.Time
	Timein  time.Time
	Timeout time.Time
	Hours   string
	Notes   sql.NullString
}

func (q *Queries) AddTracking(ctx context.Context, arg AddTrackingParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTracking,
		arg.Userid,
		arg.Date,
		arg.Timein,
		arg.Timeout,
		arg.Hours,
		arg.Notes,
	)
}

const addUser = `-- name: AddUser :execresult
INSERT INTO users (
  RoleID, Email, Password, FirstName, LastName, OfficeID, Birthdate, Active
) VALUES (
  1, ?, ?, ?, ?, ?, ?, 1
)
`

type AddUserParams struct {
	Email     string
	Password  string
	Firstname string
	Lastname  string
	Officeid  int32
	Birthdate time.Time
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUser,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Lastname,
		arg.Officeid,
		arg.Birthdate,
	)
}

const banUser = `-- name: BanUser :exec
UPDATE users
SET Active = 0
WHERE ID = ?
`

func (q *Queries) BanUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, banUser, id)
	return err
}

const createCountry = `-- name: CreateCountry :execresult
INSERT INTO countries (
  Name
) VALUES (
  ?
)
`

func (q *Queries) CreateCountry(ctx context.Context, name string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCountry, name)
}

const deleteCountry = `-- name: DeleteCountry :exec
DELETE FROM countries
WHERE Id = ?
`

func (q *Queries) DeleteCountry(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCountry, id)
	return err
}

const deleteTracking = `-- name: DeleteTracking :exec
DELETE FROM tracking
WHERE ID = ?
`

func (q *Queries) DeleteTracking(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTracking, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE ID = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAuthToken = `-- name: GetAuthToken :one
SELECT TokenID, UserID, AuthToken FROM authentication_tokens
WHERE UserID = ? ORDER BY GeneratedAt DESC LIMIT 1
`

type GetAuthTokenRow struct {
	Tokenid   int64
	Userid    int32
	Authtoken string
}

func (q *Queries) GetAuthToken(ctx context.Context, userid int32) (GetAuthTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthToken, userid)
	var i GetAuthTokenRow
	err := row.Scan(&i.Tokenid, &i.Userid, &i.Authtoken)
	return i, err
}

const getCountry = `-- name: GetCountry :one
SELECT id, name FROM countries
WHERE ID = ? LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, id int32) (Country, error) {
	row := q.db.QueryRowContext(ctx, getCountry, id)
	var i Country
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getOffices = `-- name: GetOffices :many
SELECT id, countryid, title, phone, contact FROM offices
ORDER BY ID
`

func (q *Queries) GetOffices(ctx context.Context) ([]Office, error) {
	rows, err := q.db.QueryContext(ctx, getOffices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Office
	for rows.Next() {
		var i Office
		if err := rows.Scan(
			&i.ID,
			&i.Countryid,
			&i.Title,
			&i.Phone,
			&i.Contact,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTracking = `-- name: GetTracking :many

SELECT id, userid, date, timein, timeout, hours, notes FROM tracking
ORDER BY ID
`

// -- name: FindUserByEmail :one
// SELECT * FROM users
// WHERE Email = ? LIMIT 1;
func (q *Queries) GetTracking(ctx context.Context) ([]Tracking, error) {
	rows, err := q.db.QueryContext(ctx, getTracking)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tracking
	for rows.Next() {
		var i Tracking
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Date,
			&i.Timein,
			&i.Timeout,
			&i.Hours,
			&i.Notes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserData = `-- name: GetUserData :one
SELECT ID, RoleID, Email, Password, OfficeID, FirstName, LastName, Active FROM users
WHERE Email = ? LIMIT 1
`

type GetUserDataRow struct {
	ID        int32
	Roleid    int32
	Email     string
	Password  string
	Officeid  int32
	Firstname string
	Lastname  string
	Active    sql.NullBool
}

func (q *Queries) GetUserData(ctx context.Context, email string) (GetUserDataRow, error) {
	row := q.db.QueryRowContext(ctx, getUserData, email)
	var i GetUserDataRow
	err := row.Scan(
		&i.ID,
		&i.Roleid,
		&i.Email,
		&i.Password,
		&i.Officeid,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
	)
	return i, err
}

const getUserDataByID = `-- name: GetUserDataByID :one
SELECT ID, RoleID, Email, Password, OfficeID, FirstName, LastName, Active FROM users
WHERE ID = ? LIMIT 1
`

type GetUserDataByIDRow struct {
	ID        int32
	Roleid    int32
	Email     string
	Password  string
	Officeid  int32
	Firstname string
	Lastname  string
	Active    sql.NullBool
}

func (q *Queries) GetUserDataByID(ctx context.Context, id int32) (GetUserDataByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDataByID, id)
	var i GetUserDataByIDRow
	err := row.Scan(
		&i.ID,
		&i.Roleid,
		&i.Email,
		&i.Password,
		&i.Officeid,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
	)
	return i, err
}

const getUserToUpdate = `-- name: GetUserToUpdate :one
SELECT id, roleid, email, password, firstname, lastname, officeid, birthdate, active FROM users
WHERE ID = ? LIMIT 1
`

func (q *Queries) GetUserToUpdate(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserToUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Roleid,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		&i.Officeid,
		&i.Birthdate,
		&i.Active,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT ID, RoleID, Email, FirstName, LastName, OfficeID, Active FROM users
ORDER BY ID
`

type GetUsersRow struct {
	ID        int32
	Roleid    int32
	Email     string
	Firstname string
	Lastname  string
	Officeid  int32
	Active    sql.NullBool
}

func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Roleid,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
			&i.Officeid,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAuthToken = `-- name: InsertAuthToken :exec
INSERT INTO authentication_tokens (
  UserID, AuthToken, GeneratedAt, ExpiresAt
) VALUES (
  ?, ?, ?, ?
)
`

type InsertAuthTokenParams struct {
	Userid      int32
	Authtoken   string
	Generatedat time.Time
	Expiresat   time.Time
}

func (q *Queries) InsertAuthToken(ctx context.Context, arg InsertAuthTokenParams) error {
	_, err := q.db.ExecContext(ctx, insertAuthToken,
		arg.Userid,
		arg.Authtoken,
		arg.Generatedat,
		arg.Expiresat,
	)
	return err
}

const listCountries = `-- name: ListCountries :many
SELECT id, name FROM countries
ORDER BY Name
`

func (q *Queries) ListCountries(ctx context.Context) ([]Country, error) {
	rows, err := q.db.QueryContext(ctx, listCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Country
	for rows.Next() {
		var i Country
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unbanUser = `-- name: UnbanUser :exec
UPDATE users
SET Active = 1
WHERE ID = ?
`

func (q *Queries) UnbanUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, unbanUser, id)
	return err
}

const updateTracking = `-- name: UpdateTracking :exec
UPDATE tracking
SET
  Date = ?,
  TimeIn = ?,
  TimeOut = ?,
  Hours = ?,
  Notes = ?
WHERE ID = ?
`

type UpdateTrackingParams struct {
	Date    time.Time
	Timein  time.Time
	Timeout time.Time
	Hours   string
	Notes   sql.NullString
	ID      int32
}

func (q *Queries) UpdateTracking(ctx context.Context, arg UpdateTrackingParams) error {
	_, err := q.db.ExecContext(ctx, updateTracking,
		arg.Date,
		arg.Timein,
		arg.Timeout,
		arg.Hours,
		arg.Notes,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
  RoleID = ?,
  Email = ?,
  FirstName = ?,
  LastName = ?,
  OfficeID = ?
WHERE ID = ?
`

type UpdateUserParams struct {
	Roleid    int32
	Email     string
	Firstname string
	Lastname  string
	Officeid  int32
	ID        int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Roleid,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Officeid,
		arg.ID,
	)
	return err
}
