package models

import "github.com/igor-sn/user-register/src/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()
	sql := `DELETE FROM users WHERE id = ?`
	res, err := conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
