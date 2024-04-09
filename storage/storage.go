package storage

import (
	"chinaski-magazin/storage/structs"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9"
	"log"
)

func DB() *goqu.Database {
	conn, err := sql.Open("sqlite3", "identifier.sqlite")
	db := goqu.New("sqlite3", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SelectToShelf() ([]structs.Shelf, error) {
	var shelfs []structs.Shelf
	query := DB().Select(structs.Shelf{}).From("shelf")
	if err := query.ScanStructs(&shelfs); err != nil {
		return nil, fmt.Errorf("error selecting to shelf: %w", err)
	}
	return shelfs, nil
}

func SelectToProduct() ([]structs.Product, error) {
	var products []structs.Product
	query := DB().Select(structs.Product{}).From("products")
	if err := query.ScanStructs(&products); err != nil {
		return nil, fmt.Errorf("error selecting to product: %w", err)
	}
	return products, nil
}

func SelectToShelfProduct() ([]structs.ShelfProduct, error) {
	var shelfProducts []structs.ShelfProduct
	query := DB().Select(structs.ShelfProduct{}).From("shelf_products")
	if err := query.ScanStructs(&shelfProducts); err != nil {
		return nil, fmt.Errorf("error selecting to shelf products: %w", err)
	}
	return shelfProducts, nil
}

func SelectToOrderProducts() ([]structs.OrderProduct, error) {
	var ordersProducts []structs.OrderProduct
	query := DB().Select(structs.OrderProduct{}).From("order_products")
	if err := query.ScanStructs(&ordersProducts); err != nil {
		return nil, fmt.Errorf("error selecting to order products: %w", err)
	}
	return ordersProducts, nil
}
