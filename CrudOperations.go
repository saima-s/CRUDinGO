package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type customer struct {
	ID   int
	Name string
	DOB  string
	Age  int
	Addr address
}
type address struct {
	ID           int
	City         string
	State        string
	StreetNumber string
	CustId       int
}

//type Service struct {
//	db *sql.DB
//}

func getDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:saima@123Sult@/CustomerDB?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetCustomerByName(w http.ResponseWriter, r *http.Request) {
	// gets slice of customers
	db := getDBConnection()
	param, ok := r.URL.Query()["name"]
	query := "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId;"
	var data []interface{}
	if ok && len(param) > 0 {
		query = "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.Name = ?;"
		data = append(data, param[0])
	}
	rows, err := db.Query(query, data...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var result []customer
	for rows.Next() {
		var c customer
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
			log.Fatal(err)
		}
		result = append(result, c)
	}

	byte, _ := json.Marshal(result)

	_, err = w.Write(byte)
	if err != nil {
		panic(err.Error())
	}
}
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err1 := strconv.Atoi(param["id"])
	fmt.Println("err1 came")
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]customer(nil))
		if err != nil {
			panic(err.Error())
		}
	} else {
		var ids []interface{}
		ids = append(ids, id)
		fmt.Println(id)
		db := getDBConnection()
		query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
		rows, err := db.Query(query, ids...)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var c customer
		for rows.Next() {
			if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
				log.Fatal(err)
			}
		}
		err = json.NewEncoder(w).Encode(c)
		if err != nil {
			panic(err.Error())
		}
	}
}
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]customer(nil))
		if err != nil {
			panic(err.Error())
		}
	} else {

		db := getDBConnection()
		var c customer
		var cust []interface{}
		err2 := json.Unmarshal(body, &c)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode([]customer(nil))
			if err != nil {
				panic(err.Error())
			}
		} else {
			cust = append(cust, c.Name)
			cust = append(cust, c.DOB)
			cust = append(cust, c.Age)
			if c.Age <= 18 {
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
				c.Addr.CustId = int(id)
				fmt.Println(c)
				byte, _ := json.Marshal(c)
				_, err = w.Write(byte)
				if err != nil {
					panic(err.Error())
				}
			} else {
				err := json.NewEncoder(w).Encode(customer{})
				if err != nil {
					panic(err.Error())
				}
			}
		}
	}
}
func EditCustomerDetails(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]customer(nil))
		if err != nil {
			panic(err.Error())
		}
	} else {
		db := getDBConnection()
		var c customer
		err := json.Unmarshal(body, &c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode([]customer(nil))
			if err != nil {
				panic(err.Error())
			}
			//fmt.Println("in err ")
			//log.Fatal(err)
		} else {
			param := mux.Vars(r)
			id, err1 := strconv.Atoi(param["id"])
			if err1 != nil {
				w.WriteHeader(http.StatusBadRequest)
				err := json.NewEncoder(w).Encode([]customer(nil))
				if err != nil {
					panic(err.Error())
				}
			} else {
				if c.Name != "" {
					_, err := db.Exec("update Customers set Name=? where ID=?", c.Name, id)
					if err != nil {
						panic(err.Error())
						err := json.NewEncoder(w).Encode(customer{})
						if err != nil {
							err.Error()
						}
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
				query = query[:len(query)-1]
				query += "where CustId = ? and ID = ?"
				data = append(data, id)
				data = append(data, c.Addr.ID)
				_, err = db.Exec(query, data...)

				if err != nil {
					log.Fatal(err)
				}
				err = json.NewEncoder(w).Encode(c)
				if err != nil {
					panic(err.Error())
				}
			}
		}
	}
}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	var ids []interface{}
	param := mux.Vars(r)
	id, err1 := strconv.Atoi(param["id"])
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]customer(nil))
		if err != nil {
			panic(err.Error())
		}
	}
	ids = append(ids, id)
	db := getDBConnection()
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	if !rows.Next() {
		fmt.Println("No rows")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]customer(nil))
	} else {
		query = `DELETE  FROM Customers where ID =?; `
		_, err1 = db.Exec(query, ids...)
		if err1 != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var c customer
		for rows.Next() {
			if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Age, &c.Addr.ID, &c.Addr.City, &c.Addr.State, &c.Addr.StreetNumber, &c.Addr.CustId); err != nil {
				log.Fatal(err)
			}
		}
		json.NewEncoder(w).Encode(c)
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customer", GetCustomerByName).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", GetCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/customer", CreateCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", EditCustomerDetails).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", DeleteCustomerById).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8070", r))
}
