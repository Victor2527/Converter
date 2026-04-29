package main

import (
	"fmt"
)

const gallonToLiter = 3.78541

func getPriceRU() float64 {
	return 68.0
}
func getPriceUS() float64 {
	return 4.0
}

type Text struct {
	title     string
	priceRU   string
	priceUS   string
	menu1     string
	menu2     string
	choose    string
	litersIn  string
	gallonsIn string
	rubCost   string
	usdCost   string
	invalid   string
}

func getText(lang int) Text {
	if lang == 2 {
		return Text{
			title:     "Fuel prices:",
			priceRU:   "Approx price: %.2f RUB/liter",
			priceUS:   "Approx price: %.2f USD/gallon",
			menu1:     "1 - Liters -> Gallons",
			menu2:     "2 - Gallons -> Liters",
			choose:    "Choose: ",
			litersIn:  "Enter liters: ",
			gallonsIn: "Enter gallons: ",
			rubCost:   "Approx cost: %.2f RUB",
			usdCost:   "Approx cost: %.2f USD",
			invalid:   "Invalid choice",
		}
	}

	return Text{
		title:     "Цены на бензин:",
		priceRU:   "Приблизительная стоимость: %.2f RUB/литр",
		priceUS:   "Приблизительная стоимость: %.2f USD/галлон",
		menu1:     "1 - Литры -> Галлоны",
		menu2:     "2 - Галлоны -> Литры",
		choose:    "Выбор: ",
		litersIn:  "Введите литры: ",
		gallonsIn: "Введите галлоны: ",
		rubCost:   "Приблизительная стоимость: %.2f RUB",
		usdCost:   "Приблизительная стоимость: %.2f USD",
		invalid:   "Неверный выбор",
	}
}

func main() {
	var lang int
	var choice int
	var amount float64

	fmt.Println("Select language:")
	fmt.Println("1 - RUS")
	fmt.Println("2 - ENG")
	fmt.Print("Choice: ")
	fmt.Scan(&lang)

	if lang != 1 && lang != 2 {
		fmt.Println("Invalid language selection")
		return
	}

	t := getText(lang)

	priceRU := getPriceRU()
	priceUS := getPriceUS()

	fmt.Println("\n" + t.title)
	fmt.Printf(t.priceRU+"\n", priceRU)
	fmt.Printf(t.priceUS+"\n\n", priceUS)

	fmt.Println(t.menu1)
	fmt.Println(t.menu2)
	fmt.Print(t.choose)
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Print(t.litersIn)
		fmt.Scan(&amount)

		gallons := amount / gallonToLiter

		costRU := amount * priceRU
		costUS := gallons * priceUS

		fmt.Printf("\n%.2f liters = %.2f gallons\n", amount, gallons)
		fmt.Printf(t.rubCost+"\n", costRU)
		fmt.Printf(t.usdCost+"\n", costUS)

	case 2:
		fmt.Print(t.gallonsIn)
		fmt.Scan(&amount)

		liters := amount * gallonToLiter

		costUS := amount * priceUS
		costRU := liters * priceRU

		fmt.Printf("\n%.2f gallons = %.2f liters\n", amount, liters)
		fmt.Printf(t.usdCost+"\n", costUS)
		fmt.Printf(t.rubCost+"\n", costRU)

	default:
		fmt.Println(t.invalid)
	}
}
