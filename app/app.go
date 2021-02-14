package app

import (
	"log"
	"net/http"

	"github.com/enidovasku/banking/domain"
	"github.com/enidovasku/banking/service"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
