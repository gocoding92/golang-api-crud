package controllers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go-echo-api/config"
	"net/http"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FetchListPegawaiController(c echo.Context) error {
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

	query := "SELECT * FROM pegawai ORDER BY id ASC"

	rows, err := db.Query(query)
	defer rows.Close()

	var arrobj []Pegawai
	var obj Pegawai

	for rows.Next() {
		err := rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)

		if err != nil {
			panic(err)
		}

		arrobj = append(arrobj, obj)
	}

	res := &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    arrobj,
	}

	return c.JSON(http.StatusOK, res)
}
