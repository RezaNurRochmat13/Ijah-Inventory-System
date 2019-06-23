package model

type ListAllPenjualan struct {
	IDPenjualan    string
	KodePesan      string
	NamaBarang     string
	NamaUkuran     string
	NamaWarna      string
	TanggalJual    string
	JumlahOrder    int
	HargaJual      int
	TotalPenjualan int
	HargaBeli      int
	Laba           int
}
