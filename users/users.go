package users

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db, err = sql.Open("sqlite3", os.Getenv("GOPATH")+"/src/github.com/yusufemon/project3/project3.db")

type Data struct {
	Id int
	Name string
	Balance int
}

func Get() []Data {
	checkErr(err)

	rows, err := db.Query("select id, name, balance from users")
	checkErr(err)
	defer rows.Close()

	var data = []Data{}
	for rows.Next() {
		var id int
		var name string
		var balance int
		err = rows.Scan(&id, &name, &balance)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Data{id, name, balance})
	}

	return data
}

func Insert(id int, name string, balance int) string {
	checkErr(err)

	stmt, err := db.Prepare("insert into users(id, name, balance) values(?,?,?)")
	checkErr(err)
	defer stmt.Close()

	stmt.Exec(id, name, balance)
	checkErr(err)

	return "Insert Success"
}

func Update(id int, name string, balance int) string {
	checkErr(err)

	stmt, err := db.Prepare("update users set name=?, balance=? where id=?")
	checkErr(err)
	defer stmt.Close()

	stmt.Exec(name, balance, id)
	checkErr(err)

	return "Update Success"
}

func Delete(id int) string {
	checkErr(err)

	stmt, err := db.Prepare("delete from users where id=?")
	checkErr(err)
	defer stmt.Close()

	stmt.Exec(id)
	checkErr(err)

	return "Delete Success"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}