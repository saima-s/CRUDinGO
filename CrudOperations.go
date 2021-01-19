package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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
	db := getDBConnection()
	param, ok:=r.URL.Query()["name"]
	query:="SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId;"
	var data []interface{}
	if ok && len(param)>0{
		query = "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.Name = ?;"
		data = append(data, param[0])
	}
	rows,err := db.Query(query,data...)
	if err !=nil{
		panic(err.Error())
	}
	defer rows.Close()
	var result []customer
	for rows.Next() {
		var c customer
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB,&c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
			log.Fatal(err)
		}
		result = append(result, c)
		//fmt.Println(result)
	}

	byte, _:= json.Marshal(result)
	w.Write(byte)
}
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	// gets slice of customer
	var ids []interface{}
	param:=mux.Vars(r)
	ids = append(ids,param["id"])
	fmt.Println(param["id"])
	db := getDBConnection()
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows,err := db.Query(query, ids...)
	if err !=nil{
		panic(err.Error())
	}
	defer rows.Close()
	var c customer
	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB,&c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode(c)
	//byte, err:= json.Marshal(result)
	//if err!=nil{
	//	panic(err.Error())
	//}
	//w.Write(byte)
}
func CreateCustomer(w http.ResponseWriter, r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)
	db := getDBConnection()
	var c customer
	var cust []interface{}
	json.Unmarshal(body,&c)
	cust = append(cust, c.Name)
	cust = append(cust, c.DOB)
	cust = append(cust, c.Age)
	if c.Age <=18 {
		query := `INSERT INTO Customers(name, DOB, Age) VALUES(?,?,?);`
		rows, err := db.Exec(query, cust...)
		if err != nil {
			panic(err.Error())
		}
		id, _ := rows.LastInsertId()
		var addr []interface{}
		addr = append(addr, c.Addr.City)
		addr = append(addr, c.Addr.State)
		addr = append(addr, c.Addr.StreetNumber)
		addr = append(addr, id)
		query1 := `INSERT INTO Address(City,State,StreetNumber,CustId) VALUES(?,?,?,?)`
		row, err1 := db.Exec(query1, addr...)
		if err1 != nil {
			panic(err.Error())
		}
		idAddr, _ := row.LastInsertId()
		c.ID = int(id)
		c.Addr.ID = int(idAddr)
		fmt.Println(c)
		byte, _ := json.Marshal(c)
		w.Write(byte)
	}else{
		json.NewEncoder(w).Encode(customer{})
	}

}
func EditCustomerDetails(w http.ResponseWriter, r *http.Request)  {
	body, _ := ioutil.ReadAll(r.Body)
	db := getDBConnection()
	var c customer
	err:=json.Unmarshal(body,&c)
	if err!=nil{
		fmt.Println("in err ")
		log.Fatal(err)
	}
	param:=mux.Vars(r)
	id:= param["id"]
	if c.Name != "" {
		_, err := db.Exec("update Customers set Name=? where ID=?", c.Name, id)
		if err != nil {
			panic(err.Error())
			json.NewEncoder(w).Encode(customer{})
		}
	}
		var data []interface{}
		query := "update Address set "
	    if c.Addr.City != "" {
		query += "City = ? ,"
		data = append(data, c.Addr.City)
	    }
		if c.Addr.State != "" {
			query += "State = ? ,"
			data = append(data, c.Addr.State)
		}
		if c.Addr.StreetNumber != "" {
			query += "StreetNumber = ? ,"
			data = append(data, c.Addr.StreetNumber)
		}
		query=query[:len(query)-1]
		query += "where CustId = ? and ID = ?"
		data = append(data, id)
		data = append(data, c.Addr.ID)
		_, err = db.Exec(query, data...)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(c)

}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	    var ids []interface{}
		param:=mux.Vars(r)
		ids = append(ids,param["id"])
		db := getDBConnection()
		query :=`SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
		rows, err := db.Query(query, ids...)
		if err!=nil{
			panic(err.Error())
		}
		query = `DELETE  FROM Customers where ID =?; `
		_,err1 := db.Exec(query, ids...)
		if err1 !=nil{
			panic(err.Error())
		}
		defer rows.Close()
		//var result []customer
	var c customer
		for rows.Next() {

			if err := rows.Scan(&c.ID, &c.Name, &c.DOB,&c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
				log.Fatal(err)
			}
		}
	json.NewEncoder(w).Encode(c)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customer", GetCustomerByName).Methods("GET")
	r.HandleFunc("/customer/{id}", GetCustomerById).Methods("GET")
	r.HandleFunc("/customer", CreateCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", EditCustomerDetails).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", DeleteCustomerById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8070", r))
}


