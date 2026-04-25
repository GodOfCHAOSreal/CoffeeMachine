package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (m *Machine) SaveStats(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "Orders=%d\n", m.Orders)
	fmt.Fprintf(file, "Revenue=%d\n", m.Revenue)
	return nil
}

func (m *Machine) LoadStats(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	value := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, "=", 2)

		qty, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		value = append(value, qty)
	}
	m.Orders = value[0]
	m.Revenue = value[1]
	return nil
}
