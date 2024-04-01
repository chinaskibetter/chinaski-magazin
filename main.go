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

	rows, err := db.Query("SELECT s.name AS shelf_name, p.name AS product_name, p.id AS product_id, ps.quantity, op.id_order AS order_id FROM shelf s JOIN shelf_products ps ON s.id = ps.id_shelf JOIN products p ON ps.id_product = p.id JOIN order_products op ON p.id = op.id_product AND op.id_order")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var nameShelf string
	var nameProduct string
	var productID, orderID int
	var quantity int
	var lastShelfName string
	fmt.Println("=+=+==")
	fmt.Println("Страница сборки заказов", strings.Join(orderIDs, ","))

	for rows.Next() {
		err = rows.Scan(&nameShelf, &nameProduct, &productID, &quantity, &orderID)
		if err != nil {
			log.Fatal(err)
		}

		if nameShelf != lastShelfName {
			lastShelfName = nameShelf
			fmt.Println("===", nameShelf)
		}
		fmt.Printf("%s (id=%d)\n", nameProduct, productID)
		fmt.Printf("заказ %d, %d шт\n", orderID, quantity)
		fmt.Println()
	}
}
