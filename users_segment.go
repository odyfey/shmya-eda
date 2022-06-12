package main

import (
	"golang.org/x/exp/slices"
)

const (
	lastDayJanuaryHolidays = 9
	controlOrderPrice      = 800
)

func isJanuaryHolidays(day int) bool {
	return day <= lastDayJanuaryHolidays
}

func usersSegment(orders []Order) []int {
	var result []int

	for _, order := range orders {
		if order.Cutlery > cutleryControlQuantity && !isJanuaryHolidays(order.Date.Day()) && order.Price > controlOrderPrice {
			if !slices.Contains(result, order.UserID) {
				result = append(result, order.UserID)
			}
		}
	}

	return result
}
