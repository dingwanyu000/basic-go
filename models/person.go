package models

import (
	db "github.com/testGo/databases"
	"log"
)

//定义person类型结构
type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
}

func (t *Person) AddPerson() (id int64, err error) {
	var maxId = 0
	db.SqlDB.QueryRow("select id from person order by id desc").Scan(&maxId)
	rs, err := db.SqlDB.Exec("INSERT INTO person(id,first_name, last_name,user_name,pass_word) VALUES (?, '', '', ?, ?)", maxId+1, t.UserName, t.PassWord)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (t *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, user_name, pass_word FROM person")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.UserName, &person.PassWord)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) GetPerson() (person Person, err error) {
	err = db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET user_name=?, pass_word=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.UserName, p.PassWord, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
