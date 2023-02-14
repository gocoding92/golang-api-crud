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

	rows, err := db.Query("SELECT * FROM pegawai ORDER BY id ASC")

	defer rows.Close()

	for rows.Next() {
		var id int
		var nama, alamat, telepon string
		err := rows.Scan(&id, &nama, &alamat, &telepon)

		if err != nil {
			panic(err)
		}

		u := &Pegawai{
			Id:      id,
			Nama:    nama,
			Alamat:  alamat,
			Telepon: telepon,
		}

		return c.JSON(http.StatusOK, u)
	}

	return c.JSON(http.StatusOK, rows)
}
