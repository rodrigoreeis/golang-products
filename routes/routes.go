package routes

import (
	"golang-products-api/controllers"
	"net/http"
)

func Start() {
	http.HandleFunc("/", controllers.HandleHome)
	http.HandleFunc("/new-product", controllers.HandleAddProduct)
	http.HandleFunc("/insert", controllers.HandleInsertProduct)
}
