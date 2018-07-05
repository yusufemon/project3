package users

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	db   *sql.DB
	Data Data
}

type Data struct {
	Id      int
	Name    string
	Balance int
}

func New() (*User, error) {
	db, err := sql.Open("sqlite3", os.Getenv("GOPATH")+"/src/github.com/yusufemon/project3/project3.db")
	if err != nil {
		return nil, err
	}
	return &User{db: db}, nil
}

func (user User) Get() ([]Data, error) {
	db := user.db
	rows, err := db.Query("select id, name, balance from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data = []Data{}
	for rows.Next() {
		var id int
		var name string
		var balance int
		err = rows.Scan(&id, &name, &balance)
		if err != nil {
			return nil, err
		}
		data = append(data, Data{id, name, balance})
	}

	return data, nil
}

func (user User) Insert(id int, name string, balance int) error {
	db := user.db
	stmt, err := db.Prepare("insert into users(id, name, balance) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(id, name, balance)
	if err != nil {
		return err
	}

	return nil
}

func (user User) Update(id int, name string, balance int) error {
	db := user.db
	stmt, err := db.Prepare("update users set name=?, balance=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(name, balance, id)
	if err != nil {
		return err
	}

	return nil
}

func (user User) Delete(id int) error {
	db := user.db
	stmt, err := db.Prepare("delete from users where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
