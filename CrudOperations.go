package main

import (
	"database/sql"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
type customer struct {
	ID   int
	Name string
	DOB  string
	Age int
	Addr address
}
type address struct {
	ID           int
	City         string
	State        string
	StreetNumber string
	CustId       int
}
type Service struct {
	db *sql.DB
}
func getDBConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:saima@123Sult@/CustomerDB?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	return  db
}

func GetCustomerByName(w http.ResponseWriter, r *http.Request) {
	// gets slice of customers
	w.Write([]byte("Gorilla!\n"))
}
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	// gets slice of customer
	w.Write([]byte("Gorilla!\n"))
}
func CreateCustomer(w http.ResponseWriter, r *http.Request)  {
	// writes the same data that we created
	w.Write([]byte("Gorilla!\n"))
}
func EditCustomerDetails(w http.ResponseWriter, r *http.Request)  {
	// writes the same data that we modified
	w.Write([]byte("Gorilla!\n"))
}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	// returns success on finding that Id
	w.Write([]byte("Gorilla!\n"))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customer/{name:[A-Za-z]}", GetCustomerByName).Methods("GET")
	r.HandleFunc("/customer/{id:[0-9]}", GetCustomerById).Methods("GET")
	r.HandleFunc("/customer", CreateCustomer).Methods("POST")
	r.HandleFunc("/customer", EditCustomerDetails).Methods("PUT")
	r.HandleFunc("/customer/{id:[0-9]}", DeleteCustomerById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}


