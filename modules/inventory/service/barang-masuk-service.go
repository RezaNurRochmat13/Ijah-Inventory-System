package service

import (
	"database/sql"
	"log"
	"svc-inventory-go/modules/inventory/model"
	"svc-inventory-go/modules/inventory/repository"
)

func FetchAllBarangMasukService(database *sql.DB) ([]model.ListBarangMasuk, error) {

	var (
		modelListBarangMasuk model.ListBarangMasuk
		resultAllBarangMasuk []model.ListBarangMasuk
	)

	barangMasukRepository := repository.FindAllBarangMasuk()

	queryAllBarangMasuk, errorHandlerQuery := database.Query(barangMasukRepository)

	if errorHandlerQuery != nil {
		log.Fatalf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryAllBarangMasuk.Next() {
		errorHandlerScanData := queryAllBarangMasuk.Scan(
			&modelListBarangMasuk.IDBarangMasuk,
			&modelListBarangMasuk.NamaBarang,
			&modelListBarangMasuk.NamaUkuran,
			&modelListBarangMasuk.NamaWarna,
			&modelListBarangMasuk.TanggalMasuk)

		if errorHandlerScanData != nil {
			log.Fatalf("Error query exception %v", errorHandlerScanData.Error())
		}

		resultAllBarangMasuk = append(resultAllBarangMasuk, modelListBarangMasuk)
	}

	return resultAllBarangMasuk, nil
}
