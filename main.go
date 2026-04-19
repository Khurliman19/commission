package main

import (
	"fmt"
	"project/internal/commission"
)

func main() {
	var amount int64
	var isAlifInput int

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&isAlifInput)

	isAlif := isAlifInput == 1

	if !commission.Validate(amount) {
		fmt.Println("Ошибка: сумма вне допустимого диапазона")
		return
	}

	comm := commission.Calculate(amount, isAlif)
	total := amount + comm

	fmt.Println("======== Чек ========")
	fmt.Println("Услуга: SHOKIROV SHUHRAT")
	fmt.Printf("Сумма: %d сум\n", amount)
	fmt.Printf("Комиссия: %d сум\n", comm)
	fmt.Printf("Итого: %d сум\n", total)
	fmt.Println("Статус: Исполнено")
	fmt.Println("=====================")
}
