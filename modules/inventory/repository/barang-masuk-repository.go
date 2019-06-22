package repository

func FindAllBarangMasuk() string {
	return "SELECT id_barang_masuk, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"tanggal_barang_masuk FROM barang_masuk"
}

func FindBarangMasukById() string {
	return "SELECT id_barang_masuk, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"tanggal_barang_masuk, " +
		"harga_beli, " +
		"jumlah_pesan, " +
		"jumlah_diterima, " +
		"total_harga_beli, " +
		"no_kwitansi, " +
		"keterangan " +
		"FROM barang_masuk " +
		"WHERE barang_masuk.id_barang_masuk = ?"
}
