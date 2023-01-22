// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package mysql

import (
	"database/sql"
	"time"
)

type AuthenticationToken struct {
	Tokenid     int64
	Userid      int32
	Authtoken   string
	Generatedat time.Time
	Expiresat   time.Time
}

type Country struct {
	ID   int32
	Name string
}

type Office struct {
	ID        int32
	Countryid int32
	Title     string
	Phone     string
	Contact   string
}

type Role struct {
	ID    int32
	Title string
}

type Tracking struct {
	ID      int32
	Userid  int32
	Date    time.Time
	Timein  time.Time
	Timeout time.Time
	Hours   string
	Notes   sql.NullString
}

type User struct {
	ID        int32
	Roleid    int32
	Email     string
	Password  string
	Firstname string
	Lastname  string
	Officeid  sql.NullInt32
	Birthdate sql.NullTime
	Active    sql.NullBool
}
