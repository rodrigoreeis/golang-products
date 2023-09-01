package controllers

import (
	"golang-products-api/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HandleHome(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "AddProduct", nil)
}

func HandleInsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err.Error())
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			panic(err.Error())
		}

		models.InsertProducts(name, description, price, quantity)
		http.Redirect(w, r, "/", 301)
	}
}
