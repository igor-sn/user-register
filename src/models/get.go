package models

import "github.com/igor-sn/user-register/src/db"

func Get(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM users WHERE id=$1`

	row := conn.QueryRow(sql, id)

	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Email)

	return
}
