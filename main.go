package main

import (
	"golang-products-api/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.Start()
	http.ListenAndServe(":8081", nil)
}
