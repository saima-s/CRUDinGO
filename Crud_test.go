package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"

	//"database/sql"
	//"errors"
	"net/http/httptest"

	//"reflect"
	"testing"
)
//func TestGetCustomerByNameReturnsResult(t *testing.T) {
//	var testCases = []struct {
//		inp string
//		out []customer
//	}{
//
//		{"?name=CustomerA", []customer{{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10, Addr: address{1, "Hyderabad", "Telanagana", "12", 1}}}},
//		{"?name=CustomerB", []customer{{ID: 2, Name: "CustomerB", DOB: "10/10/2010",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}}},
//		{"?", []customer{customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10, Addr: address{1, "Hyderabad", "Telanagana", "12", 1}},
//			                     customer{ID: 2, Name: "CustomerB", DOB: "10/10/2010",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}}},
//		{"?name=xyz", []customer(nil)},
//	}
//	for ind := range testCases{
//		req := httptest.NewRequest("GET", "http://localhost:8080/customer" +testCases[ind].inp,nil)
//		w:=httptest.NewRecorder()
//		GetCustomerByName(w, req)
//		resp := w.Body.Bytes()
//		var cust []customer
//		json.Unmarshal(resp,&cust)
//		fmt.Println("cust is ", cust)
//		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
//			t.Errorf("Expected %T, %v Output %T, %v",testCases[ind].out, testCases[ind].out,cust, cust)
//		}
//
//	}
//	}
//func TestGetCustomerByIdReturnsResult(t *testing.T) {
//	var testCases = []struct {
//		inp string
//		out customer
//	}{
//
//		{"1", customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10,Addr: address{1, "Hyderabad", "Telanagana", "12", 1}}},
//		{"2", customer{ID: 2, Name: "CustomerB", DOB: "10/10/2010",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}},
//		{"70", customer{}},
//	}
//	for ind := range testCases{
//		req := httptest.NewRequest("GET", "http://localhost:8070/customer/",nil)
//		//req := httptest.NewRequest(http.MethodGet, "http://localhost:8070/customer/{id}", nil)
//		req = mux.SetURLVars(req, map[string]string{"id": testCases[ind].inp})
//		w:=httptest.NewRecorder()
//		GetCustomerById(w, req)
//		resp := w.Result()
//		body,err := ioutil.ReadAll(resp.Body)
//		if err != nil{
//
//		}
//		var cust customer
//		json.Unmarshal(body,&cust)
//		fmt.Printf("body is: %v",string(body))
//
//		if !reflect.DeepEqual(cust, testCases[ind].out){
//			t.Errorf("Expected %v Output %v",testCases[ind].out,cust)
//		}
//	}
//
//}
//func TestCreateCustomerReturnsResult(t *testing.T) {
//	var testCases = []struct {
//		inp customer
//		out customer
//	}{
//
//		{customer{ID: 8, Name: "CustomerA8", DOB: "10/10/2010", Age:10,Addr: address{8, "Hyderabad", "Telangana", "12", 1}}, customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Addr: address{1, "Hyderabad", "Telangana", "12", 8}}},
//		//{customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}},
//		//{customer{ID: 3, Name: "CustomerC", DOB: "10/10/2000", Age: 20, Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{}},
//	}
//
//	for ind := range testCases{
//		byte,_:= json.Marshal(testCases[ind].inp)
//		req := httptest.NewRequest("POST", "http://localhost:8070/customer",bytes.NewBuffer(byte))
//		w:=httptest.NewRecorder()
//		CreateCustomer(w, req)
//		resp := w.Body.Bytes()
//		var cust customer
//		json.Unmarshal(resp,&cust)
//		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
//			t.Errorf("Expected %v Output %v",testCases[ind].out,cust)
//		}
//	}

//}
//func TestUpdateCustomerReturnsResult(t *testing.T) {
//	var testCases = []struct {
//		inp customer
//		out customer
//	}{
//
//		{customer{ID: 1, Name: "CustomerA", Addr: address{1, "Hyderabad", "Telangana", "12001", 1}}, customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Addr: address{1, "Hyderabad", "Telangana", "12001", 1}}},
//		{customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}},
//		{customer{ID: 100, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{}},
//	}
//	for ind := range testCases{
//		byte,_:= json.Marshal(testCases[ind].inp)
//		req := httptest.NewRequest("PUT", "http://localhost:8080/customer/{id}",bytes.NewBuffer(byte))
//		w:=httptest.NewRecorder()
//		EditCustomerDetails(w, req)
//		resp := w.Body.Bytes()
//		var cust customer
//		json.Unmarshal(resp,&cust)
//		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
//			t.Errorf("Expected %v Output %v",testCases[ind].out,cust)
//		}
//	}
//}
func TestDeleteCustomerReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp string
		out customer

	}{

		{"9", customer{ID: 9, Name: "CustomerA", DOB: "10/10/2010", Age: 10,Addr: address{9, "Hyderabad", "Telanagana", "12", 9}}},
		//{"2", customer{ID: 2, Name: "CustomerB", DOB: "10/10/2010", Age: 10, Addr: address{2, "Patna", "Bihar", "121", 2}}},
		//{"86", customer{}},
	}
	for ind := range testCases{
		req := httptest.NewRequest(http.MethodDelete, "http://localhost:8070/customer/",nil)
		//req := httptest.NewRequest(http.MethodGet, "http://localhost:8070/customer/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": testCases[ind].inp})
		w:=httptest.NewRecorder()
		DeleteCustomerById(w, req)
		resp := w.Body.Bytes()
		var cust customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
			t.Errorf("Expected %v Output %v",testCases[ind].out,cust)
		}
	}
}