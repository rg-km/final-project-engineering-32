package repository

type User struct {
	ID       int64  `db:"id_user"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type Buku struct {
	ID          int64  `db:"id_buku"`
	Pengarang   string `db:"pengarang"`
	JudulBuku   string `db:"judul_buku"`
	TahunTerbit int    `db:"tahun_terbit"`
	NoBuku      string `db:"no_buku"`
	ListBuku    string `db:"list_buku"`
}

type Peminjaman struct {
	ID             int64  `db:"id_pinjam"`
	IdBuku         int64  `db:"id_buku"`
	IdUser         int64  `db:"id_user"`
	TanggalPinjam  string `db:"tgl_pinjam"`
	TanggalKembali string `db:"tgl_kembali"`
}

// type Sales struct {
// 	ID          int64     `db:"id"`
// 	Date        time.Time `db:"date"`
// 	Category    string    `db:"category"`
// 	ProductID   int64     `db:"product_id"`
// 	ProductName string    `db:"product_name"`
// 	Price       int       `db:"price"`
// 	Quantity    int       `db:"quantity"`
// 	Total       int       `db:"total"`
// }

// type GetSalesRequest struct {
// 	StartPeriod *time.Time `db:"start_period"`
// 	EndPeriod   *time.Time `db:"end_period"`
// 	ProductID   int64      `db:"product_id"`
// 	ProductName string     `db:"product_name"`
// }

func (r *Peminjaman) IsEmptyRequest() bool {
	if r.IdBuku == 0 && r.IdUser == 0 && r.TanggalPinjam == "" {
		return true
	}

	return false
}

func (r *Peminjaman) IsValidRequest() bool {

	return r.IdBuku != 0 && r.IdUser != 0 && r.TanggalPinjam != "" && r.TanggalKembali != ""
}
