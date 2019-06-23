package repository

func FindAllPenjualan() string {
	return "SELECT id_penjualan, " +
		"kode_pesan, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"tanggal_jual, " +
		"jumlah_order, " +
		"harga_jual, " +
		"total_penjualan, " +
		"harga_beli, " +
		"laba FROM penjualan"
}
