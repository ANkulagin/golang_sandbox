package store

// Интерфейс для хранения данных о клиентах
type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

// Конкретная реализация интерфейса CustomerStorage
type MySQLStore struct{}

func (ms *MySQLStore) StoreCustomer(customer Customer) error {
	// Реализация сохранения клиента в MySQL
	return nil
}

func (ms *MySQLStore) GetCustomer(id string) (Customer, error) {
	// Реализация получения клиента из MySQL
	return Customer{}, nil
}

func (ms *MySQLStore) GetAllCustomers() ([]Customer, error) {
	// Реализация получения всех клиентов из MySQL
	return nil, nil
}

func (ms *MySQLStore) GetCustomersWithoutContract() ([]Customer, error) {
	// Реализация получения клиентов без контракта из MySQL
	return nil, nil
}

func (ms *MySQLStore) GetCustomersWithNegativeBalance() ([]Customer, error) {
	// Реализация получения клиентов с отрицательным балансом из MySQL
	return nil, nil
}

// Структура Customer
type Customer struct {
	ID       string
	Balance  float64
	Contract bool
}
