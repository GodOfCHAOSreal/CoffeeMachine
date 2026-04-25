package domain

type Storage interface {
	GetAll() map[string]int
	Add(name string, qty int) error
	Set(name string, qty int) error
	HasStock(drinkName string, recipes map[string]Recipe) bool
	UseForDrink(drinkName string, recipes map[string]Recipe)
}
