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
		no_buku
	FROM  buku`
	rows, err := bk.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var buku Buku
		err := rows.Scan(&buku.ID, &buku.Pengarang, &buku.JudulBuku, &buku.TahunTerbit, &buku.NoBuku)
		if err != nil {
			return nil, err
		}
		books = append(books, buku)
	}
	return books, nil
	//endanswer return []CartItem{}, nil
}

// func (c *Buku) buku(productID int64) (buku, error) {
// 	var Buku buku
// 	var sqlStatement string
// 	//TODO : you must fetch the cart by product id
// 	//HINT : you can use the where statement
// 	//beginanswer
// 	sqlStatement = `
// 	SELECT
// 	c.id_buku,
// 	c.pengarang,
// 	c.judul_buku,
// 	p.tahun_terbit,
// 	p.no_buku
// 	FROM cart_items c
// 	INNER JOIN products p
// 	ON c.product_id = p.id
// 	WHERE c.product_id = ?
// 	`

// 	err := c.db.QueryRow(sqlStatement, productID).Scan(&cartItem.ID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.Category, &cartItem.ProductName, &cartItem.Price)
// 	if err != nil {
// 		return cartItem, err
// 	}

// 	return cartItem, nil
// 	//endanswer return CartItem{}, nil
// }

// func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
// 	// TODO: you must insert the cart item
// 	//beginanswer
// 	sqlStatement := `
// 	INSERT INTO cart_items (product_id, quantity)
// 	VALUES (?, ?)
// 	`
// 	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// 	//endanswer return nil
// }

// func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
// 	//TODO : you must update the quantity of the cart item
// 	//beginanswer
// 	sqlStatement := `
// 	UPDATE cart_items
// 	SET quantity = quantity + 1
// 	WHERE id = ?
// 	`
// 	_, err := c.db.Exec(sqlStatement, cartItem.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// 	//endanswer return nil
// }

// func (c *CartItemRepository) ResetCartItems() error {
// 	//TODO : you must reset the cart items
// 	//HINT : you can use the delete statement
// 	//beginanswer
// 	sqlStatement := `
// 	DELETE FROM cart_items
// 	`
// 	_, err := c.db.Exec(sqlStatement)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// 	//endanswer return nil
// }

// func (c *CartItemRepository) TotalPrice() (int, error) {
// 	var sqlStatement string
// 	//TODO : you must calculate the total price of the cart items
// 	//HINT : you can use the sum statement
// 	//beginanswer
// 	sqlStatement = `
// 	SELECT
// 	    SUM(p.price * c.quantity)
// 	FROM cart_items c
// 	INNER JOIN products p
// 	ON c.product_id = p.id
// `
// 	var totalPrice int
// 	err := c.db.QueryRow(sqlStatement).Scan(&totalPrice)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return totalPrice, nil
// 	//endanswer return 0, nil
// }
