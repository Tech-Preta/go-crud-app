package main

import (
	"go-crud-app/db"
	"go-crud-app/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	db.InitDatabase()

	// Carregar templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	router := mux.NewRouter()

	// Servir arquivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rotas
	router.HandleFunc("/", ListItems).Methods("GET")
	router.HandleFunc("/create", ShowCreateForm).Methods("GET")
	router.HandleFunc("/create", CreateItem).Methods("POST")
	router.HandleFunc("/edit/{id:[0-9]+}", ShowEditForm).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditItem).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", ShowDeleteForm).Methods("GET")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteItem).Methods("POST")

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func ListItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	db.DB.Find(&items)
	templates.ExecuteTemplate(w, "index.html", items)
}

func ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create.html", nil)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	item := models.Item{Name: name, Price: price}
	db.DB.Create(&item)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowEditForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var item models.Item
	db.DB.First(&item, id)
	templates.ExecuteTemplate(w, "edit.html", item)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	name := r.FormValue("name")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	var item models.Item
	db.DB.First(&item, id)
	item.Name = name
	item.Price = price
	db.DB.Save(&item)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowDeleteForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var item models.Item
	db.DB.First(&item, id)
	templates.ExecuteTemplate(w, "delete.html", item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var item models.Item
	db.DB.Delete(&item, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
