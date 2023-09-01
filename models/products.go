package models

import (
	"golang-products-api/config"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := config.ConnectionDB()
	res, err := db.Query("SELECT * FROM products")

	if err != err {
		panic(err.Error())
	}

	var products []Product

	for res.Next() {
		product := Product{}
		err := res.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func InsertProducts(name string, description string, price float64, quantity int) {
	db := config.ConnectionDB()
	stmt, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(name, description, price, quantity)

	defer db.Close()
}
