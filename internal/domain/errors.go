package domain

import "errors"

var (
	ErrorDrinkNotFound        = errors.New("напиток не найден")
	ErrorNotEnoughPay         = errors.New("недостаточно оплаты")
	ErrorNotEnoughIngridients = errors.New("не хватает ингридиентов")
	ErrorNotCorrect           = errors.New("некорректные параметры")
)
