package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Customer struct (Model) ...
type Suppliers struct {
	SupplierID   string `json:"SupplierID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Region       string `json:"Region"`
	PostalCode   string `json:"PostalCode"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	Fax          string `json:"Fax"`
	HomePage     string `json:"HomePage"`
}

// Get all orders

func getSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var suppliers []Suppliers

	sql := `SELECT
				SupplierID,
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(Fax,'') Fax,
				IFNULL(HomePage,'') HomePage
			FROM suppliers`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var supplier Suppliers
		err := result.Scan(&supplier.SupplierID, &supplier.CompanyName, &supplier.ContactName,
			&supplier.ContactTitle, &supplier.Address, &supplier.City, &supplier.Region, &supplier.PostalCode, &supplier.Country,
			&supplier.Phone, &supplier.Fax, &supplier.HomePage)

		if err != nil {
			panic(err.Error())
		}
		suppliers = append(suppliers, supplier)
	}

	json.NewEncoder(w).Encode(suppliers)
}

func createSuppliers(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		SupplierID := r.FormValue("SupplierID")
		CompanyName := r.FormValue("CompanyName")
		ContactName := r.FormValue("ContactName")
		ContactTitle := r.FormValue("ContactTitle")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		Phone := r.FormValue("Phone")
		Fax := r.FormValue("Fax")
		HomePage := r.FormValue("HomePage")

		stmt, err := db.Prepare("INSERT INTO suppliers (SupplierID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax,HomePage) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(SupplierID, CompanyName, ContactName, ContactTitle, Address, City, Region, PostalCode, Country, Phone, Fax, HomePage)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}

func getSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var suppliers []Suppliers
	params := mux.Vars(r)

	sql := `SELECT
				SupplierID,
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(Fax,'') Fax,
				IFNULL(HomePage,'') HomePage
			FROM suppliers WHERE SupplierID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var supplier Suppliers

	for result.Next() {

		err := result.Scan(&supplier.SupplierID, &supplier.CompanyName, &supplier.ContactName,
			&supplier.ContactTitle, &supplier.Address, &supplier.City, &supplier.Region, &supplier.PostalCode, &supplier.Country,
			&supplier.Phone, &supplier.Fax, &supplier.HomePage)

		if err != nil {
			panic(err.Error())
		}

		suppliers = append(suppliers, supplier)
	}

	json.NewEncoder(w).Encode(suppliers)
}

func updateSUppliers(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newCompanyName := r.FormValue("CompanyName")

		stmt, err := db.Prepare("UPDATE suppliers SET CompanyName = ? WHERE SupplierID = ?")

		_, err = stmt.Exec(newCompanyName, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Customer with CustomerID = %s was updated", params["id"])
	}
}

func deleteSuppliers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM suppliers WHERE SupplierID = ?")

	_, err = stmt.Exec(params["id"])

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "Supplier with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var suppliers []Suppliers

	SupplierID := r.FormValue("SupplierID")
	CompanyName := r.FormValue("CompanyName")

	sql := `SELECT
				SupplierID,
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(Fax,'') Fax,
				IFNULL(HomePage,'') HomePage
			FROM suppliers WHERE SupplierID = ? AND CompanyName = ?`

	result, err := db.Query(sql, SupplierID, CompanyName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var supplier Suppliers

	for result.Next() {

		err := result.Scan(&supplier.SupplierID, &supplier.CompanyName, &supplier.ContactName,
			&supplier.ContactTitle, &supplier.Address, &supplier.City, &supplier.Region, &supplier.PostalCode, &supplier.Country,
			&supplier.Phone, &supplier.Fax, &supplier.HomePage)

		if err != nil {
			panic(err.Error())
		}

		suppliers = append(suppliers, supplier)
	}

	json.NewEncoder(w).Encode(suppliers)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/suppliers", getSupplier).Methods("GET")
	r.HandleFunc("/suppliers/{id}", getSuppliers).Methods("GET")
	r.HandleFunc("/suppliers", createSuppliers).Methods("POST")
	r.HandleFunc("/suppliers/{id}", updateSUppliers).Methods("PUT")
	r.HandleFunc("/suppliers/{id}", deleteSuppliers).Methods("DELETE")

	//New
	r.HandleFunc("/getsuppliers", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
