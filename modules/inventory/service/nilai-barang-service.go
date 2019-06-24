package service

import (
	"database/sql"
	"log"
	"svc-inventory-go/modules/inventory/model"
	"svc-inventory-go/modules/inventory/repository"
)

func FetchTotalLaba(database *sql.DB) ([]model.NilaiBarang, error) {
	var (
		modelNilaiBarang  model.NilaiBarang
		resultNilaiBarang []model.NilaiBarang
	)

	totalLabaRepository := repository.FindAllTotalLaba()

	queryTotalLaba, errorHandlerQuery := database.Query(totalLabaRepository)

	if errorHandlerQuery != nil {
		log.Printf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryTotalLaba.Next() {
		errorHandlerScanData := queryTotalLaba.Scan(
			&modelNilaiBarang.TotalPenjualan,
			&modelNilaiBarang.LabaTotal)

		if errorHandlerScanData != nil {
			log.Printf("Error query exception %v", errorHandlerQuery.Error())
		}

		resultNilaiBarang = append(resultNilaiBarang, modelNilaiBarang)
	}

	return resultNilaiBarang, nil

}
