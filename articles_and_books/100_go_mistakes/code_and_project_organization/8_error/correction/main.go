package store

// Структура Customer
type Customer struct {
	ID      string
	Balance float64
}

// Структура Contract
type Contract struct {
	ID     string
	Active bool
}

// Структура Store
type Store struct{}

// GetCustomer Метод GetCustomer возвращает конкретный тип Customer
func (s *Store) GetCustomer(id string) (Customer, error) {
	// Логика получения клиента
	return Customer{}, nil
}

// SetCustomer Метод SetCustomer принимает конкретный тип Customer
func (s *Store) SetCustomer(id string, customer Customer) error {
	// Логика установки клиента
	return nil
}

// GetContract Метод GetContract возвращает конкретный тип Contract
func (s *Store) GetContract(id string) (Contract, error) {
	// Логика получения контракта
	return Contract{}, nil
}

// Метод SetContract принимает конкретный тип Contract
func (s *Store) SetContract(id string, contract Contract) error {
	// Логика установки контракта
	return nil
}
