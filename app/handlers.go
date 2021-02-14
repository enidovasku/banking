package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/enidovasku/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
