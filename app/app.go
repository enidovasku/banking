package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/enidovasku/banking/domain"
	"github.com/enidovasku/banking/service"

	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Env var nor defined")
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	//Init Db client
	dbClient := getDbClient()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
