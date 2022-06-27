package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type PeminjamanRepository struct {
	db *sql.DB
}

func NewPeminjamanRepository(db *sql.DB) *PeminjamanRepository {
	return &PeminjamanRepository{db: db}
}

func (bk *PeminjamanRepository) GetListPeminjaman() ([]Peminjaman, error) {
	var sqlStatement string
	var Lending []Peminjaman

	//TODO: add sql statement here
	//HINT: join table cart_items and products

	//beginanswer
	sqlStatement = `
	SELECT
		id_pinjam,
		id_buku,
		id_user,
		tgl_pinjam,
		tgl_kembali
	FROM  peminjaman`
	rows, err := bk.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var peminjaman Peminjaman
		err := rows.Scan(&peminjaman.ID, &peminjaman.IdBuku, &peminjaman.IdUser, &peminjaman.TanggalPinjam, &peminjaman.TanggalKembali)
		if err != nil {
			return nil, err
		}
		Lending = append(Lending, peminjaman)
	}
	return Lending, nil
	//endanswer return []CartItem{}, nil

}

func (u *PeminjamanRepository) InsertPeminjaman(id_buku int64, id_user int64, tgl_pinjam string, tgl_kembali string) (int64, error) {
	//beginanswer
	res, err := u.db.Exec("INSERT INTO peminjaman (id_buku, id_user, tgl_pinjam, tgl_kembali) VALUES (?, ?, ?, ?)",
		id_buku, id_user, tgl_pinjam, tgl_kembali)
	if err != nil {
		return 0, err
	}
	id_pinjam, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id_pinjam, nil
	//endanswer return nil
}
