package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "backend/db/perpustakaan.db")
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
	CONSTRAINT peminjaman_FK FOREIGN KEY (id_user) REFERENCES "user"(id_user) ON DELETE CASCADE,
	CONSTRAINT peminjaman_FK_1 FOREIGN KEY (id_buku) REFERENCES buku(id_buku) ON DELETE CASCADE
);

// CREATE TABLE IF NOT EXISTS sales (
//     id integer not null primary key AUTOINCREMENT,
//     date DATE not null,
//     product_id integer not null,
//     quantity integer not null,
//     FOREIGN KEY (product_id) REFERENCES products(id)
// );

// INSERT INTO users(username, password, role, loggedin) VALUES
//     ('aditira', '1234', 'admin', false),
//     ('dina', '4321', 'employee', false),
//     ('dito', '2552', 'employee', false);

INSERT INTO products(product_name, category, price, quantity) VALUES
    ('Orange', 'Fruits', 5000, 100),
    ('Apple', 'Fruits', 2000, 100),
    ('Melon', 'Fruits', 4000, 100),
    ('Watermelon', 'Fruits', 10000, 100),
    ('Banana', 'Fruits', 4000, 100),
    ('Carrot', 'Vegetables', 2000, 100),
    ('Broccoli', 'Vegetables', 5200, 100),
    ('Cucumber', 'Vegetables', 3400, 100),
    ('Potato', 'Vegetables', 6500, 100),
    ('Tomato', 'Vegetables', 2200, 100),
    ('Coffee', 'Drink', 4300, 100),
    ('Milk', 'Drink', 4000, 100),
    ('Tea', 'Drink', 2700, 100);`)

	if err != nil {
		panic(err)
	}
}
