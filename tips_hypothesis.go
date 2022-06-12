package main

import (
	"log"
)

const cutleryControlQuantity = 2

func tipsHypothesisForHolidays(orders []Order) bool {
	log.Println("january holidays tips hypothesis")

	var tipsWithMoreCutlery, tipsWithFewerCutlery int
	var ordersWithoutTips, ordersWithMoreCutlery, ordersWithFewerCutlery int

	for _, order := range orders {
		if isJanuaryHolidays(order.Date.Day()) {
			if order.Cutlery > cutleryControlQuantity {
				if order.Tips == 0 {
					ordersWithoutTips++
				} else {
					tipsWithMoreCutlery += order.Tips
					ordersWithMoreCutlery++
				}
			} else {
				tipsWithFewerCutlery += order.Tips
				ordersWithFewerCutlery++
			}
		}
	}

	log.Printf(
		"ordersWithoutTips = %d, ordersWithMoreCutlery = %d, ordersWithFewerCutlery = %d",
		ordersWithoutTips, ordersWithMoreCutlery, ordersWithFewerCutlery,
	)

	log.Printf(
		"tipsWithMoreCutlery = %d, tipsWithFewerCutlery = %d",
		tipsWithMoreCutlery, tipsWithFewerCutlery,
	)

	averageTipsWithMoreCutlery := tipsWithMoreCutlery / (ordersWithMoreCutlery + ordersWithoutTips)
	averageTipsWithFewerCultery := tipsWithFewerCutlery / ordersWithFewerCutlery

	log.Printf(
		"average tips for orders:\n\twith more than two cutlery = %d\n\twith less than two cutlery = %d\n",
		averageTipsWithMoreCutlery, averageTipsWithFewerCultery,
	)

	return averageTipsWithMoreCutlery > averageTipsWithFewerCultery
}

func tipsHypothesisForAnotherDays(orders []Order) bool {
	log.Println("another days tips hypothesis")

	var tipsWithMoreCutlery, tipsWithFewerCutlery int
	var ordersWithoutTips, ordersWithMoreCutlery, ordersWithFewerCutlery int

	for _, order := range orders {
		if !isJanuaryHolidays(order.Date.Day()) {
			if order.Cutlery > cutleryControlQuantity {
				if order.Tips == 0 {
					ordersWithoutTips++
				} else {
					tipsWithMoreCutlery += order.Tips
					ordersWithMoreCutlery++
				}
			} else {
				tipsWithFewerCutlery += order.Tips
				ordersWithFewerCutlery++
			}
		}
	}

	log.Printf(
		"ordersWithoutTips = %d, ordersWithMoreCutlery = %d, ordersWithFewerCutlery = %d",
		ordersWithoutTips, ordersWithMoreCutlery, ordersWithFewerCutlery,
	)

	log.Printf(
		"tipsWithMoreCutlery = %d, tipsWithFewerCutlery = %d",
		tipsWithMoreCutlery, tipsWithFewerCutlery,
	)

	averageTipsWithMoreCutlery := tipsWithMoreCutlery / (ordersWithMoreCutlery + ordersWithoutTips)
	averageTipsWithFewerCultery := tipsWithFewerCutlery / ordersWithFewerCutlery

	log.Printf(
		"average tips for orders:\n\twith more than two cutlery = %d\n\twith less than two cutlery = %d\n",
		averageTipsWithMoreCutlery, averageTipsWithFewerCultery,
	)

	return averageTipsWithMoreCutlery > averageTipsWithFewerCultery
}
