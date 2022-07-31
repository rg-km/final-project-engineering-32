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

func (r *Peminjaman) IsEmptyRequest() bool {
	if r.IdBuku == 0 && r.IdUser == 0 && r.TanggalPinjam == "" {
		return true
	}

	return false
}

func (r *Peminjaman) IsValidRequest() bool {

	return r.IdBuku != 0 && r.IdUser != 0 && r.TanggalPinjam != "" && r.TanggalKembali != ""
}
