package repository

type user struct {
	ID       int64  `db:"id_user"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Loggedin bool   `db:"loggedin"`
	Token    string `db:"token"`
}

type buku struct {
	ID          int64  `db:"id_buku"`
	Category    string `db:"pengarang"`
	ProductName string `db:"judul_buku"`
	Price       int    `db:"tahun_terbit"`
	Quantity    int    `db:"no_buku"`
}

type peminjaman struct {
	ID          int64  `db:"id_pinjam"`
	Category    string `db:"id_buku"`
	ProductID   int64  `db:"id_user"`
	ProductName string `db:"tgl_pinjam"`
	Price       int    `db:"price"`
	Quantity    int    `db:"quantity"`
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

func (r GetSalesRequest) IsEmptyRequest() bool {
	if r.StartPeriod == nil && r.EndPeriod == nil && r.ProductName == "" {
		return true
	}

	return false
}

func (r GetSalesRequest) IsValidRequest() bool {
	if r.StartPeriod == nil && r.EndPeriod != nil {
		return false
	}
	if r.StartPeriod != nil && r.EndPeriod == nil {
		return false
	}

	return true
}
