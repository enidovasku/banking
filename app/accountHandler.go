package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/enidovasku/banking/dto"

	"github.com/enidovasku/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, account)
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	//Get account_id and customer_id from the URL
	vars := mux.Vars(r)
	account_id := vars["account_id"]
	customer_id := vars["customer_id"]

	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.AccountId = account_id
	request.CustomerId = customer_id

	//make transaction
	account, appError := h.service.MakeTransaction(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, account)
}
