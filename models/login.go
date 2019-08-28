package models

import (
	"database/sql"
	db "github.com/testGo/databases"
)

type Login struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (p *Login) Login(login Login) bool {
	var name sql.NullString
	db.SqlDB.QueryRow("select user_name from person where user_name = ? and pass_word = ?", login.UserName, login.PassWord).Scan(&name)
	if name.Valid {
		return true
	} else {
		return false
	}
}
