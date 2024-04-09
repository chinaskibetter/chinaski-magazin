package main

import (
	"bufio"
	"chinaski-magazin/storage"
	"chinaski-magazin/storage/structs"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	orderIDs := getUserInput()

	products, err := storage.SelectToProduct()
	if err != nil {
		log.Fatal(err)
	}

	shelfProducts, err := storage.SelectToShelfProduct()
	if err != nil {
		log.Fatal(err)
	}
	ordersProducts, err := storage.SelectToOrderProducts()
	if err != nil {
		log.Fatal(err)
	}
	shelfs, err := storage.SelectToShelf()
	if err != nil {
		log.Fatal(err)
	}

	printProductsOnShelves(orderIDs, products, shelfProducts, ordersProducts, shelfs)
}

func getUserInput() []int {
	fmt.Println("Введите номера заказов через запятую:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	orderIDsStr := strings.Split(scanner.Text(), ",")
	var orderIDs []int
	for _, id := range orderIDsStr {
		num, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		orderIDs = append(orderIDs, num)
	}
	return orderIDs
}

func printProductsOnShelves(orderIDs []int, products []structs.Product, shelfProducts []structs.ShelfProduct, ordersProducts []structs.OrderProduct, shelfs []structs.Shelf) {
	fmt.Println("=+=+=+=")
	fmt.Printf("Страница сборки заказов %v\n", orderIDs)

	for _, shelf := range shelfs {
		fmt.Printf("\n===%s\n", shelf.Name)
		for _, product := range products {
			for _, orderProduct := range ordersProducts {
				for _, shelfProduct := range shelfProducts {
					for _, orderID := range orderIDs {
						if int64(orderID) == orderProduct.IdOrder && orderProduct.IdProduct == product.Id && product.Id == shelfProduct.IdProduct && shelfProduct.IdShelf == shelf.Id {
							fmt.Printf("%s (id=%d)\n", product.Name, product.Id)
							fmt.Printf("заказ %d, %d шт\n\n", orderProduct.IdOrder, shelfProduct.Quantity)
						}
					}
				}
			}
		}
	}
}
