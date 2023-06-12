package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "A", City: "P", Pincode: "1"},
		{Id: "2", Name: "B", City: "Q", Pincode: "2"},
		{Id: "3", Name: "C", City: "R", Pincode: "3"},
	}
	return CustomerRepositoryStub{customers: customers}
}
