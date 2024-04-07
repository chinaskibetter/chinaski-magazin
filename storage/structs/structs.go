package structs

import (
	_ "github.com/doug-martin/goqu/v9"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

var Products []Product

type ShelfProduct struct {
	IdProduct int64 `db:"id_product"`
	IdShelf   int64 `db:"id_shelf"`
	Quantity  int64 `db:"quantity"`
	IsMain    int64 `db:"is_main"`
}

var ShelfProducts []ShelfProduct

type Shelf struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

var Shelfs []Shelf

type OrderProduct struct {
	IdOrder   int64 `db:"id_order"`
	IdProduct int64 `db:"id_product"`
}

var OrdersProducts []OrderProduct
