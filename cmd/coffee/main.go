package main

import (
	"bufio"
	"coffee/internal/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	stock   = domain.NewStock()
	machine = domain.NewMachine(stock)
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) < 1 {
			fmt.Println("usage")
		}
		switch args[0] {
		case "help":
			fmt.Println("usage")
		case "menu":
			for drink, recipe := range domain.DefaultRecipes() {
				fmt.Printf("%s - %d\n", drink, recipe.Price)
			}
		case "stock":
			if len(args) < 2 {
				fmt.Println("usage")
				continue
			}
			switch args[1] {
			case "get":
				for item, qty := range stock.GetAll() {
					fmt.Printf("%s=%d\n", item, qty)
				}
			case "add":
				if len(args) < 4 {
					fmt.Println("usage")
					continue
				}
				qty, err := strconv.Atoi(args[3])
				if err != nil {
					fmt.Println(domain.ErrorNotCorrect)
					continue
				}
				err = stock.Add(args[2], qty)
				if err != nil {
					fmt.Println(err)
					continue
				}
			case "set":
				if len(args) < 4 {
					fmt.Println("usage")
					continue
				}
				qty, err := strconv.Atoi(args[3])
				if err != nil {
					fmt.Println(domain.ErrorNotCorrect)
					continue
				}
				err = stock.Set(args[2], qty)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		case "brew":
			if len(args) < 4 {
				fmt.Println("usage")
				continue
			}
			payment, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Println(domain.ErrorNotCorrect)
			}
			steps, err := machine.Brew(args[1], payment)
			if err != nil {
				fmt.Println(err)
			}
			for _, step := range steps {
				fmt.Println(step)
			}
			fmt.Printf("Ваш %s готов!\n", args[1])
		default:
			fmt.Println("usage")

		}
	}
}
