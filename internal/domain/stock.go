package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	if qty < 0 {
		return ErrorNotCorrect
	}
	s.items[name] = qty
	return nil
}

func (s *Stock) HasStock(drinkName string, recipes map[string]Recipe) bool {
	for item, qty := range recipes[drinkName].Ingridients {
		if qty > s.items[item] {
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

func (s *Stock) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for name, qty := range s.items {
		fmt.Fprintf(file, "%s=%d\n", name, qty)
	}
	return nil
}

func LoadStock(path string) (*Stock, error) {
	file, err := os.Open(path)
	if err != nil {
		return NewStock(), nil
	}
	defer file.Close()

	stock := NewStock()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		qty, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		fmt.Println(name, qty)
		err = stock.Set(name, qty)
		if err != nil {
			return nil, err
		}
	}
	return stock, nil
}
