package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id 			int
	name 		string
	inventory 	int
	price 		int
}

func main() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product: ", p.id, " ", p.name, " ", p.inventory, " ", p.price)
	}
}