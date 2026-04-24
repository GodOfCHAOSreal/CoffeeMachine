package domain

type Stock struct {
	items map[string]int
}

func NewStock() *Stock {
	return &Stock{
		items: map[string]int{
			"beans": 0,
			"water": 0,
			"milk":  0,
			"sugar": 0,
		},
	}
}

func (s *Stock) GetAll() map[string]int {
	return s.items
}

func (s *Stock) Add(name string, qty int) error {
	if _, ok := s.items[name]; !ok {
		return ErrorNotCorrect
	}
	if qty < 1 {
		return ErrorNotCorrect
	}
	s.items[name] += qty
	return nil
}

func (s *Stock) Set(name string, qty int) error {
	if _, ok := s.items[name]; !ok {
		return ErrorNotCorrect
	}
	if qty <= 0 {
		return ErrorNotCorrect
	}
	s.items[name] = qty
	return nil
}

func (s *Stock) HasStock(drinkName string, recipes map[string]Recipe) bool {
	for item, qty := range recipes[drinkName].Ingridients {
		if qty <= s.items[item] {
			return false
		}
	}
	return true
}

func (s *Stock) UseForDrink(drinkName string, recipes map[string]Recipe) {
	for item, qty := range recipes[drinkName].Ingridients {
		s.items[item] -= qty
	}
}
