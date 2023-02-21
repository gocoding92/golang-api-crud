package models

import (
	"database/sql"
	"fmt"
	"go-echo-api/config"
	"go-echo-api/database"
	"net/http"
)

type Pegawai struct {
	Id      int `json:"id"`
	Nama    int `json:"nama"`
	Alamat  int `json:"alamat"`
	Telepon int `json:"telepon"`
}

func FetchListPegawaiModel() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func FetchCreatePegawaiModel(nama string, alamat string, telepon string) (Response, error) {
	var res Response
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `INSERT INTO pegawai (nama, alamat, telepon) VALUES ($1, $2, $3) RETURNING id`

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(nama, alamat, telepon)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{}

	return res, nil
}

func FetchUpdatePegawaiModel(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `UPDATE pegawai SET nama = $1, alamat = $2, telepon = $3 WHERE id = $4  `

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(nama, alamat, telepon, id)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{}

	return res, nil
}

func FetchDeletePegawaiModel(id int) (Response, error) {
	var res Response
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `DELETE FROM pegawai WHERE id = $1  `

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(id)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{}

	return res, nil
}
