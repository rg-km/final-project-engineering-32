package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "./db/perpustakaan.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS user (
    id_user integer not null primary key AUTOINCREMENT,
    username text(255) not null,
    password text(255) not null,
    role varchar(255) not null
);

CREATE TABLE IF NOT EXISTS buku (
    id_buku integer not null primary key AUTOINCREMENT,
    pengarang text(255) not null,
    judul_buku text(255) not null,
    tahun_terbit integer not null,
    no_buku text not null
);

CREATE TABLE IF NOT EXISTS peminjaman (
	id_pinjam INTEGER PRIMARY KEY AUTOINCREMENT,
	id_buku INTEGER,
	id_user INTEGER,
	tgl_pinjam TEXT,
    tgl_kembali TEXT,
	CONSTRAINT peminjaman_FK FOREIGN KEY (id_user) REFERENCES "user"(id_user) ON DELETE CASCADE,
	CONSTRAINT peminjaman_FK_1 FOREIGN KEY (id_buku) REFERENCES buku(id_buku) ON DELETE CASCADE
);`)

	if err != nil {
		panic(err)
	}
}
