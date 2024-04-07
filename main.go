package main

import (
	"bufio"
	"chinaski-magazin/storage"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

func main() {
	storage.ConnectToDB()
	storage.SelectToOrderProducts()
	storage.SelectToProduct()
	storage.SelectToShelf()
	storage.SelectToShelfProduct()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	orderIDs := strings.Split(scanner.Text(), ",")

	fmt.Println("=+=+==")
	fmt.Println("Страница сборки заказов", strings.Join(orderIDs, ","))

}
