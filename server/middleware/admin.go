package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/misha1350/mospolytech-web-app/db/mysql"
)

func GetUsers(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	queries := mysql.New(db)

	users, err := queries.GetUsers(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var usersMap []map[string]interface{}
	for _, user := range users {
		userMap := map[string]interface{}{
			"id":        user.ID,
			"role":      user.Roleid,
			"email":     user.Email,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"office":    user.Officeid,
			"active":    user.Active,
		}
		usersMap = append(usersMap, userMap)
	}

	c.JSON(200, gin.H{"users": usersMap})
}

// package mysql

// import (
// 	"context"
// 	"time"
// )

// type GetUsersRow struct {
// 	ID        int64     `db:"id"`
// 	Email     string    `db:"email"`
// 	Firstname string    `db:"firstname"`
// 	Lastname  string    `db:"lastname"`
// 	CreatedAt time.Time `db:"created_at"`
// 	UpdatedAt time.Time `db:"updated_at"`
// 	Birthdate time.Time `db:"birthdate"`
// 	Office    string    `db:"office"`
// 	Active    bool      `db:"active"`
// }

// func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
// 	const query = `
// 		SELECT *
// 		FROM users
// 	`
// 	rows, err := q.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []GetUsersRow
// 	for rows.Next() {
// 		var user GetUsersRow
// 		if err := rows.Scan(
// 			&user.ID,
// 			&user.Email,
// 			&user.Firstname,
// 			&user.Lastname,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 			&user.Birthdate,
// 			&user.Office,
// 			&user.Active,
// 		); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
