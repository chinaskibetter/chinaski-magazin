package structs

import (
	_ "github.com/doug-martin/goqu/v9"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type ShelfProduct struct {
	IdProduct int64 `db:"id_product"`
	IdShelf   int64 `db:"id_shelf"`
	Quantity  int64 `db:"quantity"`
	IsMain    bool  `db:"is_main"`
}

type Shelf struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type OrderProduct struct {
	IdOrder   int64 `db:"id_order"`
	IdProduct int64 `db:"id_product"`
}
