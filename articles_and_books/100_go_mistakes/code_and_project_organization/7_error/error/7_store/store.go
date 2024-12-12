package __store

import "fmt"

// Интерфейс для хранения данных о клиентах
type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

// Конкретная реализация интерфейса CustomerStorage
type InMemoryStore struct {
	customers map[string]Customer
}

func NewInMemoryStore() CustomerStorage {
	return &InMemoryStore{
		customers: make(map[string]Customer),
	}
}

func (s *InMemoryStore) StoreCustomer(customer Customer) error {
	s.customers[customer.ID] = customer
	return nil
}

func (s *InMemoryStore) GetCustomer(id string) (Customer, error) {
	customer, exists := s.customers[id]
	if !exists {
		return Customer{}, fmt.Errorf("customer with ID %s not found", id)
	}
	return customer, nil
}

func (s *InMemoryStore) GetAllCustomers() ([]Customer, error) {
	var list []Customer
	for _, customer := range s.customers {
		list = append(list, customer)
	}
	return list, nil
}

func (s *InMemoryStore) GetCustomersWithoutContract() ([]Customer, error) {
	var list []Customer
	for _, customer := range s.customers {
		if !customer.Contract {
			list = append(list, customer)
		}
	}
	return list, nil
}

func (s *InMemoryStore) GetCustomersWithNegativeBalance() ([]Customer, error) {
	var list []Customer
	for _, customer := range s.customers {
		if customer.Balance < 0 {
			list = append(list, customer)
		}
	}
	return list, nil
}

// Структура Customer
type Customer struct {
	ID       string
	Balance  float64
	Contract bool
}
