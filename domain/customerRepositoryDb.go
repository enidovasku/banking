package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/enidovasku/banking/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if status != "" {
		findAllSql += " where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	} else {
		rows, err = d.client.Query(findAllSql)
	}

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning rows " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}

		log.Println("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
