package domain

import "errors"

var (
	ErrorDrinkNotFound        = errors.New("Напиток не найден! Такое мы здесь не подаём, Друг!")
	ErrorNotEnoughPay         = errors.New("Денег, увы, недостаточно:(")
	ErrorNotEnoughIngridients = errors.New("На складе не хватает ингридиентов для этого напитка!")
	ErrorNotCorrect           = errors.New("Ты ошибся параметром, Друг!")
)
