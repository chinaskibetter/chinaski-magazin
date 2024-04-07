package storage

import (
	"chinaski-magazin/storage/structs"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

func ConnectToDB() *sql.DB {
	conn, err := sql.Open("sqlite3", "identifier.sqlite")
	if err != nil {
		panic(err)
	}
	return conn
}

func SelectToShelf() {
	db := goqu.New("sqlite3", ConnectToDB())
	query := db.Select("id", "name").From("shelf")
	if err := query.ScanStructs(&structs.Shelfs); err != nil {
		fmt.Println(err.Error())
		return
	}
	// [{1 Стеллаж А} {2 Стеллаж Б} {3 Стеллаж В} {4 Стеллаж З} {5 Стеллаж Ж}]
}

func SelectToProduct() {
	db := goqu.New("sqlite3", ConnectToDB())
	query := db.Select("id", "name").From("products")
	if err := query.ScanStructs(&structs.Products); err != nil {
		fmt.Println(err.Error())
		return
	}
	// [{1 Ноутбук} {2 Телевизор} {3 Телефон} {4 Системный блок} {5 Часы} {6 Микрофон}]
}

func SelectToShelfProduct() {
	db := goqu.New("sqlite3", ConnectToDB())
	query := db.Select("id_product", "id_shelf", "quantity", "is_main").From("shelf_products")
	if err := query.ScanStructs(&structs.ShelfProducts); err != nil {
		fmt.Println(err.Error())
		return
	}
	// [{1 1 2 0} {2 1 3 0} {1 1 3 0} {3 2 1 1} {4 5 4 0} {5 5 1 1} {6 5 1 0}]
}

func SelectToOrderProducts() {
	db := goqu.New("sqlite3", ConnectToDB())
	query := db.Select("id_order", "id_product").
		From("order_products").
		Join(goqu.T("products"),
			goqu.On(goqu.T("products").Col("id").Eq(goqu.T("order_products").Col("id_product"))))

	if err := query.ScanStructs(&structs.OrdersProducts); err != nil {
		fmt.Println(err.Error())
		return
	}
	// [{10 1} {10 3} {10 6} {11 2} {14 1} {14 4} {15 5}]
}
