package service

import (
	"database/sql"
	"log"
	"svc-inventory-go/modules/inventory/model"
	"svc-inventory-go/modules/inventory/repository"
)

func FetchAllStokBarang(database *sql.DB) ([]model.StokBarang, error) {
	var (
		modelStokBarang  model.StokBarang
		resultStokBarang []model.StokBarang
	)

	stokBarangRepository := repository.FindAllStokBarang()

	queryStokBarang, errorHandlerQuery := database.Query(stokBarangRepository)

	if errorHandlerQuery != nil {
		log.Printf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryStokBarang.Next() {
		errorHandlerScanData := queryStokBarang.Scan(
			&modelStokBarang.IDBarang,
			&modelStokBarang.KodeBarang,
			&modelStokBarang.NamaBarang,
			&modelStokBarang.NamaUkuran,
			&modelStokBarang.NamaWarna,
			&modelStokBarang.JumlahStokSekarang)

		if errorHandlerScanData != nil {
			log.Printf("Error query exception %v", errorHandlerQuery.Error())
		}

		resultStokBarang = append(resultStokBarang, modelStokBarang)
	}

	return resultStokBarang, nil
}
