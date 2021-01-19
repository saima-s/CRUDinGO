package main

import (
	"bytes"
	"encoding/json"
	"reflect"

	"net/http"

	//"database/sql"
	//"errors"
	"net/http/httptest"
	//"reflect"

	//"reflect"
	"testing"
)
func TestGetCustomerByNameReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp string
		out []customer
	}{

		{"CustomerA", []customer{{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10, Addr: address{1, "Hyderabad", "Telangana", "12", 1}}}},
		{"CustomerB", []customer{{ID: 2, Name: "CustomerB", DOB: "10/10/2011",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}}},
		{"", []customer{customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10, Addr: address{1, "Hyderabad", "Telangana", "12", 1}},
			                     customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}}},
		{"xyz", []customer(nil)},
	}
	for ind := range testCases{
		req := httptest.NewRequest("GET", "http://localhost:8080/customer/" +testCases[ind].inp,nil)
		w:=httptest.NewRecorder()
		GetCustomerByName(w, req)
		resp := w.Body.Bytes()
		var cust []customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
			t.Errorf("FAILED ")
		}

	}
	}
func TestGetCustomerByIdReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp string
		out customer
	}{

		{"1", customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age: 10,Addr: address{1, "Hyderabad", "Telangana", "12", 1}}},
		{"2", customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011",Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}},
		{"70", customer(nil)},
	}
	for ind := range testCases{
		req := httptest.NewRequest("GET", "http://localhost:8080/customer/" +testCases[ind].inp,nil)
		w:=httptest.NewRecorder()
		GetCustomerById(w, req)
		resp := w.Body.Bytes()
		var cust customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
			t.Errorf("FAILED ")
		}
	}

}
func TestCreateCustomerReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp customer
		out customer
	}{

		{customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Age:10,Addr: address{1, "Hyderabad", "Telangana", "12", 1}}, customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Addr: address{1, "Hyderabad", "Telangana", "12", 1}}},
		{customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Age:10, Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}},
		{customer{ID: 3, Name: "CustomerC", DOB: "10/10/2000", Age: 20, Addr: address{2, "Patna", "Bihar", "121", 2}}, customer(nil)},
	}

	for ind := range testCases{
		byte,_:= json.Marshal(testCases[ind].inp)
		req := httptest.NewRequest("POST", "http://localhost:8080/customer/",bytes.NewBuffer(byte))
		w:=httptest.NewRecorder()
		CreateCustomer(w, req)
		resp := w.Body.Bytes()
		var cust customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
				t.Errorf("FAILED ")
		}
	}

}
func TestUpdateCustomerReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp customer
		out customer
	}{

		{customer{ID: 1, Name: "CustomerA", Addr: address{1, "Hyderabad", "Telangana", "12001", 1}}, customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Addr: address{1, "Hyderabad", "Telangana", "12001", 1}}},
		{customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}, customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}},
		{customer{ID: 100, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}, customer(nil)},
	}
	for ind := range testCases{
		byte,_:= json.Marshal(testCases[ind].inp)
		req := httptest.NewRequest("PUT", "http://localhost:8080/customer/",bytes.NewBuffer(byte))
		w:=httptest.NewRecorder()
		EditCustomerDetails(w, req)
		resp := w.Body.Bytes()
		var cust customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
			t.Errorf("FAILED ")
		}
	}
}
func TestDeleteCustomerReturnsResult(t *testing.T) {
	var testCases = []struct {
		inp string
		out customer

	}{

		{"1", customer{ID: 1, Name: "CustomerA", DOB: "10/10/2010", Addr: address{1, "Hyderabad", "Telangana", "12", 1}}},
		{"2", customer{ID: 2, Name: "CustomerB", DOB: "10/10/2011", Addr: address{2, "Patna", "Bihar", "121", 2}}},
		{"86", customer(nil)},
	}
	for ind := range testCases{
		req := httptest.NewRequest("GET", "http://localhost:8080/customer/" +testCases[ind].inp,nil)
		w:=httptest.NewRecorder()
		GetCustomerById(w, req)
		resp := w.Body.Bytes()
		var cust customer
		json.Unmarshal(resp,&cust)
		if w.Code != http.StatusOK || !reflect.DeepEqual(cust, testCases[ind].out){
			t.Errorf("FAILED ")
		}
	}
}