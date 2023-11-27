package models

import "github.com/igor-sn/user-register/src/db"

func Update(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `UPDATE users SET name = $1, age = $2, email = $3 WHERE id = $4`

	res, err := conn.Exec(sql, user.Name, user.Age, user.Email, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
