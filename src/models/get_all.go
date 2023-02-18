package models

import "github.com/igor-sn/user-register/src/db"

func GetAll() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM users`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return
}
