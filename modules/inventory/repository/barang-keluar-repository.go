package repository

func FindAllBarangKeluar() string {
	return "SELECT id_barang_keluar, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"jumlah_keluar FROM barang_keluar"
}

func FindBarangKeluarById() string {
	return "SELECT id_barang_keluar, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"jumlah_keluar, " +
		"harga_jual, " +
		"total_harga_barang_keluar, " +
		"catatan " +
		"FROM barang_keluar " +
		"WHERE barang_keluar.id_barang_keluar = ?"
}
