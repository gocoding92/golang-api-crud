package models

import (
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
