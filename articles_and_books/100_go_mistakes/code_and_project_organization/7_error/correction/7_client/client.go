package client

// Интерфейс, определенный на стороне потребителя
type customersGetter interface {
	GetAllCustomers() ([]__store.Customer, error)
}

// Структура Foo, использующая интерфейс customersGetter
type Foo struct {
	getter customersGetter
}

// Фабричная функция для создания Foo
func NewFoo(g customersGetter) Foo {
	return Foo{getter: g}
}

func (f Foo) Bar() {
	_, err := f.getter.GetAllCustomers()
	if err != nil {
		// Обработка ошибки
		return
	}
	// Работа с полученными клиентами
}
