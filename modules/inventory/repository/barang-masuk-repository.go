package repository

func FindAllBarangMasuk() string {
	return "SELECT id_barang_masuk, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"tanggal_barang_masuk FROM barang_masuk"
}
