package models

type StokProduct struct {
	IDStokMasuk       uint    `gorm:"primary_key" json:"id_stok_masuk"`
	IDProduk          uint    `json:"id_produk"`
	HargaBeli         float64 `json:"harga_beli"`
	JumlahMasuk       uint    `json:"jumlah_masuk"`
	TanggalMasuk      string  `json:"tanggal_masuk"`
	TanggalKadaluarsa string  `json:"tanggal_kadaluarsa"`
	Product           Product `gorm:"foreignkey:IDProduk;references:ID" json:"product"` // Foreign Key
}
