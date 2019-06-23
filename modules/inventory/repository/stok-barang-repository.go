package repository

func FindAllStokBarang() string {
	return "SELECT id_barang, " +
		"kode_barang, " +
		"nama_barang, " +
		"nama_ukuran, " +
		"nama_warna, " +
		"jumlah_stok FROM barang"
}
