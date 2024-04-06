package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite3", "identifier.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	orderIDs := strings.Split(scanner.Text(), ",")

	rowsShelf, err := db.Query("SELECT name FROM shelf")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsShelf.Close()

	rowsProducts, err := db.Query("SELECT name, id FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsProducts.Close()

	rowsShelfProducts, err := db.Query("SELECT quantity FROM shelf_products")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsShelfProducts.Close()

	rowsOrderProducts, err := db.Query("SELECT id_order FROM order_products")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsOrderProducts.Close()

	var nameShelf string
	var nameProduct string
	var productID, orderID int
	var quantity int
	var lastShelfName string
	fmt.Println("=+=+==")
	fmt.Println("Страница сборки заказов", strings.Join(orderIDs, ","))

	for rowsShelf.Next() {
		err = rowsShelf.Scan(&nameShelf)
		if err != nil {
			log.Fatal(err)
		}

		if nameShelf != lastShelfName {
			lastShelfName = nameShelf
			fmt.Println("===", nameShelf)
		}
	}

	for rowsProducts.Next() {
		err = rowsProducts.Scan(&nameProduct, &productID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s (id=%d)\n", nameProduct, productID)
	}

	for rowsShelfProducts.Next() {
		err = rowsShelfProducts.Scan(&quantity)
		if err != nil {
			log.Fatal(err)
		}
	}

	for rowsOrderProducts.Next() {
		err = rowsOrderProducts.Scan(&orderID)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Printf("заказ %d, %d шт\n\n", orderID, quantity)
}
