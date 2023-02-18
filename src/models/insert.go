package models

import "github.com/igor-sn/user-register/src/db"

func Insert(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO users (name, age, email) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, user.Name, user.Age, user.Email).Scan(&id)

	return
}
