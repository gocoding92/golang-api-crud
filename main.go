package main

import (
	"database/sql"
	"fmt"
	"go-echo-api/routes"

	_ "github.com/lib/pq"
)

const (
	host1     = "localhost"
	port1     = 5431
	user1     = "postgres"
	password1 = "html5123"
	dbname1   = "kepegawaian"
)

type pegawai struct {
	id      int
	nama    string
	alamat  string
	telepon string
}

func main() {

	e := routes.Init()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host1, port1, user1, password1, dbname1)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM pegawai ORDER BY id ASC")

	defer rows.Close()

	var result []pegawai

	for rows.Next() {
		var each = pegawai{}
		var err = rows.Scan(&each.id, &each.nama, &each.alamat, &each.telepon)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.nama)
	}

	e.Logger.Fatal(e.Start(":2020"))

}
