package domain

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/enidovasku/banking/logger"

	"github.com/enidovasku/banking/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if status != "" {
		findAllSql += " where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	} else {
		err = d.client.Select(&customers, findAllSql)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
