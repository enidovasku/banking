package domain

import "github.com/enidovasku/banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Enid", "Hamburg", "22301", "1995-01-01", "1"},
		{"1002", "Enido", "Hamburg", "22301", "1995-01-01", "1"},
		{"1003", "Enidov", "Hamburg", "22301", "1995-01-01", "0"},
		{"1004", "Enidova", "Hamburg", "22301", "1995-01-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
