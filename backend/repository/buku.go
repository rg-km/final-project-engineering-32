package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type BukuRepository struct {
	db *sql.DB
}

func NewBukuRepository(db *sql.DB) *BukuRepository {
	return &BukuRepository{db: db}
}

func (bk *BukuRepository) GetListBuku() ([]Buku, error) {
	var sqlStatement string
	var books []Buku

	//TODO: add sql statement here
	//HINT: join table cart_items and products

	//beginanswer
	sqlStatement = `
	SELECT
		id_buku,
		pengarang,
		judul_buku,
		tahun_terbit,
		no_buku,
		list_buku,
	FROM  buku`
	rows, err := bk.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var buku Buku
		err := rows.Scan(&buku.ID, &buku.Pengarang, &buku.JudulBuku, &buku.TahunTerbit, &buku.NoBuku, &buku.ListBuku)
		if err != nil {
			return nil, err
		}
		books = append(books, buku)
	}
	return books, nil
}

func (c *BukuRepository) AddNewBuku(pengarang string, judul_buku string, tahun_terbit int, no_buku string, list_buku string) (*string, error) {

	sqlStatement := `INSERT INTO buku (pengarang, judul_buku, tahun_terbit, no_buku, list_buku) VALUES (?, ?, ?, ?, ?)`

	_, err := c.db.Exec(sqlStatement, pengarang, judul_buku, tahun_terbit, no_buku, list_buku)
	if err != nil {
		return nil, err
	}
	return &judul_buku, nil
}

func (c *BukuRepository) DeleteBuku(id int) (bool, error) {
	sqlStatement := `DELETE from buku WHERE id_buku = ?`

	_, err := c.db.Exec(sqlStatement, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (bk *BukuRepository) CheckIfProductExists(id int) (bool, error) {
	var sqlStatement string
	var books Buku

	sqlStatement = `
	SELECT
		*
	FROM buku WHERE id_buku = ? `
	err := bk.db.QueryRow(sqlStatement, id).Scan(
		&books.ID,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}
