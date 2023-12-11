package customerscontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/cahiman99/crud_golang/libraries"
	"github.com/cahiman99/crud_golang/models"
)

var validation = libraries.NewValidation()
var customersModel = models.NewCustomersModel()

func Index(response http.ResponseWriter, request *http.Request) {
	customers, _ := customersModel.FindAll()

	data := map[string]interface{}{
		"customers": customers,
	}

	temp, err := template.ParseFiles("views/customers/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/customers/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var customers entities.customers
		customers.Name = request.Form.Get("name")
		customers.NIK = request.Form.Get("nik")
		customers.JenisKelamin = request.Form.Get("jenis_kelamin")
		customers.Tempat_lahir = request.Form.Get("tempat_lahir")
		customers.Tanggal_lahir = request.Form.Get("tanggal_lahir")
		customers.Alamat = request.Form.Get("alamat")
		customers.No_hp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customers)

		if vErrors != nil {
			data["customers"] = customers
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data customers berhasil disimpan"
			customersModel.Create(customers)
		}

		temp, _ := template.ParseFiles("views/customers/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var customers entities.customers
		customersModel.Find(id, &customers)

		data := map[string]interface{}{
			"customers": customers,
		}

		temp, err := template.ParseFiles("views/customers/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var customers entities.customers
		customers.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		customers.Name = request.Form.Get("nama_lengkap")
		customers.NIK = request.Form.Get("nik")
		customers.JenisKelamin = request.Form.Get("jenis_kelamin")
		customers.Tempat_lahir = request.Form.Get("tempat_lahir")
		customers.Tanggal_lahir = request.Form.Get("tanggal_lahir")
		customers.Alamat = request.Form.Get("alamat")
		customers.No_hp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customers)

		if vErrors != nil {
			data["customers"] = customers
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data customers berhasil diperbarui"
			customersModel.Update(customers)
		}

		temp, _ := template.ParseFiles("views/customers/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	customersModel.Delete(id)

	http.Redirect(response, request, "/customers", http.StatusSeeOther)
}
