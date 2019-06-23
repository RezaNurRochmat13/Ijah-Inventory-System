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
		log.Printf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryAllBarangMasuk.Next() {
		errorHandlerScanData := queryAllBarangMasuk.Scan(
			&modelListBarangMasuk.IDBarangMasuk,
			&modelListBarangMasuk.NamaBarang,
			&modelListBarangMasuk.NamaUkuran,
			&modelListBarangMasuk.NamaWarna,
			&modelListBarangMasuk.TanggalMasuk)

		if errorHandlerScanData != nil {
			log.Printf("Error query exception %v", errorHandlerScanData.Error())
		}

		resultAllBarangMasuk = append(resultAllBarangMasuk, modelListBarangMasuk)
	}

	return resultAllBarangMasuk, nil
}

func FetchDetailBarangMasuk(barangMasukId string, database *sql.DB) ([]model.DetailBarangMasuk, error) {
	var (
		modelDetailBarangMasuk  model.DetailBarangMasuk
		resultdetailBarangMasuk []model.DetailBarangMasuk
	)

	barangMasukDetailRepository := repository.FindBarangMasukById()

	queryDetailBarangMasuk, errorHandlerQuery := database.Query(barangMasukDetailRepository, barangMasukId)

	if errorHandlerQuery != nil {
		log.Printf("Error query exception %v", errorHandlerQuery.Error())
	}

	for queryDetailBarangMasuk.Next() {
		errorHandlerScanData := queryDetailBarangMasuk.Scan(
			&modelDetailBarangMasuk.IDBarangMasuk,
			&modelDetailBarangMasuk.NamaBarang,
			&modelDetailBarangMasuk.NamaUkuran,
			&modelDetailBarangMasuk.NamaWarna,
			&modelDetailBarangMasuk.TanggalMasuk,
			&modelDetailBarangMasuk.HargaBeli,
			&modelDetailBarangMasuk.JumlahPesan,
			&modelDetailBarangMasuk.JumlahDiterima,
			&modelDetailBarangMasuk.TotalHargaBeli,
			&modelDetailBarangMasuk.NoKwitansi,
			&modelDetailBarangMasuk.Keterangan)

		if errorHandlerScanData != nil {
			log.Printf("Error query exception %v", errorHandlerScanData.Error())
		}

		resultdetailBarangMasuk = append(resultdetailBarangMasuk, modelDetailBarangMasuk)

	}

	return resultdetailBarangMasuk, nil

}
