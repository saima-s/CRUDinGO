package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	age "github.com/bearbin/go-age"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	DOB     string  `json:"dob"`
	Address address `json:"address"`
}
type address struct {
	ID         int    `json:"id"`
	City       string `json:"city"`
	State      string `json:"state"`
	StreetName string `json:"streetName"`
	CustId     int    `json:"custId"`
}

func getAge(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

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
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Address.ID, &c.Address.City, &c.Address.State, &c.Address.StreetName, &c.Address.CustId); err != nil {
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
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode([]customer(nil))
		if err != nil {
			panic(err.Error())
		}
	} else {
		var ids []interface{}
		ids = append(ids, id)
		db := getDBConnection()
		query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
		rows, err := db.Query(query, ids...)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var c customer
		for rows.Next() {
			if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Address.ID, &c.Address.City, &c.Address.State, &c.Address.StreetName, &c.Address.CustId); err != nil {
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
			dob := c.DOB
			dob1 := strings.Split(dob, "/")
			y, _ := strconv.Atoi(dob1[2])
			m, _ := strconv.Atoi(dob1[1])
			d, _ := strconv.Atoi(dob1[0])
			getAge := getAge(y, m, d)
			if age.Age(getAge) >= 18 {
				query := `INSERT INTO Customers(name, DOB) VALUES(?,?);`
				rows, err := db.Exec(query, cust...)
				if err != nil {
					panic(err.Error())
				}
				id, _ := rows.LastInsertId()
				var addr []interface{}
				addr = append(addr, c.Address.City)
				addr = append(addr, c.Address.State)
				addr = append(addr, c.Address.StreetName)
				addr = append(addr, id)
				query1 := `INSERT INTO Address(City,State,StreetName,CustId) VALUES(?,?,?,?)`
				row, err1 := db.Exec(query1, addr...)
				if err1 != nil {
					panic(err.Error())
				}
				idAddr, _ := row.LastInsertId()
				c.ID = int(id)
				c.Address.ID = int(idAddr)
				c.Address.CustId = int(id)
				w.WriteHeader(http.StatusCreated)
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
				query1 := `SELECT * from Customers where ID =?`
				var id1 []interface{}
				id1 = append(id1, id)
				row, err := db.Query(query1, id1...)
				if err != nil {
					panic(err.Error())
				}
				if !row.Next() {
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
						var custId []interface{}
						custId = append(custId, id)
						q := `SELECT * FROM Customers INNER JOIN Address where Address.CustID =?`
						r, _ := db.Query(q, custId...)
						var cu customer
						for r.Next() {
							e := r.Scan(&cu.ID, &cu.Name, &cu.DOB, &cu.Address.ID, &cu.Address.City, &cu.Address.State, &cu.Address.StreetName, &cu.Address.CustId)
							if e != nil {
								log.Fatal(e)
							}
						}
						var data []interface{}
						query := "update Address set "
						if c.Address.City != "" {
							query += "City = ? ,"
							data = append(data, c.Address.City)
						}
						if c.Address.State != "" {
							query += "State = ? ,"
							data = append(data, c.Address.State)
						}
						if c.Address.StreetName != "" {
							query += "StreetName = ? ,"
							data = append(data, c.Address.StreetName)
						}

						query = query[:len(query)-1]
						query += "where CustId = ? and ID = ?"
						data = append(data, id)
						data = append(data, cu.Address.ID)
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
	}
}
func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	var ids []interface{}
	param := mux.Vars(r)
	id, err1 := strconv.Atoi(param["id"])
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	ids = append(ids, id)
	db := getDBConnection()
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	if !rows.Next() {
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
			if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Address.ID, &c.Address.City, &c.Address.State, &c.Address.StreetName, &c.Address.CustId); err != nil {
				log.Fatal(err)
			}
		}
		w.WriteHeader(http.StatusNoContent)
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
