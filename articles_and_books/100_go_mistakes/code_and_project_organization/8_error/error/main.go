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

// Метод Get возвращает any
func (s *Store) Get(id string) (any, error) {
	// Логика получения данных
	return nil, nil
}

// Метод Set принимает any
func (s *Store) Set(id string, v any) error {
	// Логика установки данных
	return nil
}
