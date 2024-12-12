package store

// Конкретная реализация хранилища клиентов
type MySQLStore struct{}

func (ms *MySQLStore) GetAllCustomers() ([]Customer, error) {
	// Реализация получения всех клиентов из MySQL
	return []Customer{
		{ID: "1", Balance: 100.0, Contract: true},
		{ID: "2", Balance: -50.0, Contract: false},
	}, nil
}

// Структура Customer
type Customer struct {
	ID       string
	Balance  float64
	Contract bool
}
