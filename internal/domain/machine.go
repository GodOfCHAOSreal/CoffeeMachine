package domain

type Machine struct {
	Stock   Storage
	Recipe  map[string]Recipe
	Orders  int
	Revenue int
}

func NewMachine(stock Storage) *Machine {
	return &Machine{
		Stock:  stock,
		Recipe: DefaultRecipes(),
	}
}

func (m *Machine) Brew(drinkName string, pay int) ([]string, error) {
	if _, ok := m.Recipe[drinkName]; !ok {
		return nil, ErrorDrinkNotFound
	}
	if pay < m.Recipe[drinkName].Price {
		return nil, ErrorNotEnoughPay
	}
	if !m.Stock.HasStock(drinkName, m.Recipe) {
		return nil, ErrorNotEnoughIngridients
	}
	m.Stock.UseForDrink(drinkName, m.Recipe)
	m.Orders++
	m.Revenue += m.Recipe[drinkName].Price

	return m.Recipe[drinkName].Steps, nil
}

func (m *Machine) Stats() (int, int) {
	return m.Orders, m.Revenue
}
