package model

type ListBarangMasuk struct {
	IDBarangMasuk string
	NamaBarang    string
	NamaUkuran    string
	NamaWarna     string
	TanggalMasuk  string
}

type DetailBarangMasuk struct {
	IDBarangMasuk  string
	NamaBarang     string
	NamaUkuran     string
	NamaWarna      string
	TanggalMasuk   string
	HargaBeli      int
	JumlahPesan    int
	JumlahDiterima int
	TotalHargaBeli int
	NoKwitansi     string
	Keterangan     string
}
