package service

import (
	"database/sql"
	"log"
	"svc-inventory-go/modules/inventory/model"
	"svc-inventory-go/modules/inventory/repository"
)

func FetchAllBarangKeluar(database *sql.DB) ([]model.ListBarangKeluar, error) {
	var (
		modelListBarangKeluar model.ListBarangKeluar
		resultAllBarangKeluar []model.ListBarangKeluar
	)

	barangKeluarRepository := repository.FindAllBarangKeluar()

	queryAllBarangKeluar, errorHandlerQuery := database.Query(barangKeluarRepository)

	if errorHandlerQuery != nil {
		log.Fatalf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryAllBarangKeluar.Next() {
		errorHandlerScanData := queryAllBarangKeluar.Scan(
			&modelListBarangKeluar.IDBarangKeluar,
			&modelListBarangKeluar.NamaBarang,
			&modelListBarangKeluar.NamaUkuran,
			&modelListBarangKeluar.NamaWarna,
			&modelListBarangKeluar.JumlahKeluar)

		if errorHandlerScanData != nil {
			log.Fatalf("Error query exception %v", errorHandlerScanData.Error())
		}

		resultAllBarangKeluar = append(resultAllBarangKeluar, modelListBarangKeluar)
	}
	return resultAllBarangKeluar, nil

}

func FetchDetailBarangKeluar(barangKeluarID string, database *sql.DB) ([]model.DetailBarangKeluar, error) {
	var (
		modelDetailBarangKeluar  model.DetailBarangKeluar
		resultDetailBarangKeluar []model.DetailBarangKeluar
	)

	barangKeluarRepository := repository.FindBarangKeluarById()

	queryDetailBarangKeluar, errorHandlerQuery := database.Query(barangKeluarRepository, barangKeluarID)

	if errorHandlerQuery != nil {
		log.Fatalf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryDetailBarangKeluar.Next() {
		errorHandlerScanData := queryDetailBarangKeluar.Scan(
			&modelDetailBarangKeluar.IDBarangKeluar,
			&modelDetailBarangKeluar.NamaBarang,
			&modelDetailBarangKeluar.NamaUkuran,
			&modelDetailBarangKeluar.NamaWarna,
			&modelDetailBarangKeluar.JumlahKeluar,
			&modelDetailBarangKeluar.HargaJual,
			&modelDetailBarangKeluar.TotalHargaBarangKeluar,
			&modelDetailBarangKeluar.Catatan)

		if errorHandlerScanData != nil {
			log.Fatalf("Error query exception %v", errorHandlerScanData.Error())
		}

		resultDetailBarangKeluar = append(resultDetailBarangKeluar, modelDetailBarangKeluar)

	}

	return resultDetailBarangKeluar, nil

}
