package main

import (
	"net/http"

	customerscontroller "github.com/cahiman99/go-crud/controllers/customers_conttroller"
)

func main() {
	http.HandleFunc("/", customerscontroller.Index)
	http.HandleFunc("/customers/tes", customerscontroller.Index)
	http.HandleFunc("/customers/index", customerscontroller.Index)
	http.HandleFunc("/customers/add", customerscontroller.Add)
	http.HandleFunc("/customers/edit", customerscontroller.Edit)
	http.HandleFunc("/customers/delete", customerscontroller.Delete)

	http.ListenAndServe(":3000", nil) //Untuk menjalankan server
}
