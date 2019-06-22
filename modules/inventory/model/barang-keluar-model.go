package model

type ListBarangKeluar struct {
	IDBarangKeluar string
	NamaBarang     string
	NamaUkuran     string
	NamaWarna      string
	JumlahKeluar   int
}

type DetailBarangKeluar struct {
	IDBarangKeluar         string
	NamaBarang             string
	NamaUkuran             string
	NamaWarna              string
	JumlahKeluar           int
	HargaJual              int
	TotalHargaBarangKeluar int
	Catatan                string
}
