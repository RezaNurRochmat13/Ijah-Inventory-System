package service

import (
	"database/sql"
	"log"
	"svc-inventory-go/modules/inventory/model"
	"svc-inventory-go/modules/inventory/repository"
)

func FetchAllPenjualan(database *sql.DB) ([]model.ListAllPenjualan, error) {
	var (
		modelAllPenjualan  model.ListAllPenjualan
		resultAllPenjualan []model.ListAllPenjualan
	)

	penjualanRepository := repository.FindAllPenjualan()

	queryAllPenjualan, errorHandlerQuery := database.Query(penjualanRepository)

	if errorHandlerQuery != nil {
		log.Printf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryAllPenjualan.Next() {
		errorHandlerScanData := queryAllPenjualan.Scan(
			&modelAllPenjualan.IDPenjualan,
			&modelAllPenjualan.KodePesan,
			&modelAllPenjualan.NamaBarang,
			&modelAllPenjualan.NamaUkuran,
			&modelAllPenjualan.NamaWarna,
			&modelAllPenjualan.TanggalJual,
			&modelAllPenjualan.JumlahOrder,
			&modelAllPenjualan.HargaJual,
			&modelAllPenjualan.TotalPenjualan,
			&modelAllPenjualan.HargaBeli,
			&modelAllPenjualan.Laba)

		if errorHandlerScanData != nil {
			log.Printf("Error query exception %v", errorHandlerQuery.Error())
		}

		resultAllPenjualan = append(resultAllPenjualan, modelAllPenjualan)
	}

	return resultAllPenjualan, nil
}
