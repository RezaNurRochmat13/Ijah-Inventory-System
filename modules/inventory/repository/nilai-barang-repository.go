package repository

func FindAllTotalLaba() string {
	return "SELECT total_penjualan, " +
		"SUM(total_penjualan) as laba_total FROM penjualan"
}
