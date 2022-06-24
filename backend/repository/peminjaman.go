package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type PeminjamananRepository struct {
	db *sql.DB
}

func NewPeminjamanRepository(db *sql.DB) *PeminjamananRepository {
	return &PeminjamananRepository{db: db}
}

func (bk *PeminjamananRepository) GetListPeminjaman() ([]Peminjaman, error) {
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
