package main

import (
	"fmt"
	"strconv"
)

func main() {
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var a, b string
	var op string
	var firstOp, secondOp, f int
	var isRoman bool
	var isRoman1 bool
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}
	for i := 0; i < len(roman); i++ {
		if a == roman[i] {
			firstOp = i + 1
			isRoman = true
			break
		}
	}
	for i := 0; i < len(roman); i++ {
		if b == roman[i] {
			isRoman1 = true
			secondOp = i + 1
			break
		}
	}
	if isRoman1 != isRoman {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}
	if !isRoman && !isRoman1 {
		firstOp, _ = strconv.Atoi(a)
		secondOp, _ = strconv.Atoi(b)
	}
	if (firstOp >= 1 && firstOp <= 10) || (secondOp >= 1 && secondOp <= 10) {
	} else {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — числа нужны от 1 до 10 включительно")
		return
	}
	switch op {
	case "+":
		f = firstOp + secondOp
	case "-":
		f = firstOp - secondOp
	case "*":
		f = firstOp * secondOp
	case "/":
		f = firstOp / secondOp
	default:
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}
	if isRoman == true {
		if f > 10 && f < 20 {
			ost := f % 10
			fmt.Println("X" + roman[ost-1])
		} else if f > 20 && f < 40 {
			ost := f % 10
			c := f / 10
			for i := 0; i < c; i++ {
				fmt.Print("X")
			}
			fmt.Print(roman[ost-1])
		} else if f == 40 {
			fmt.Println("XL")
		} else if f >= 41 && f <= 49 {
			ost := f - 40
			fmt.Println("XL" + roman[ost-1])
		} else if f == 50 {
			fmt.Println("L")
		} else if f >= 51 && f <= 89 {
			ost := f % 10
			c := (f - 50) / 10
			fmt.Print("L")
			for i := 0; i < c; i++ {
				fmt.Print("X")
			}
			fmt.Print(roman[ost-1])
		} else if f == 100 {
			fmt.Println("XC")
		} else if f == 90 {
			fmt.Println("XC")
		} else if f >= 91 && f <= 99 {
			ost := f - 90
			fmt.Println("XC" + roman[ost-1])
		} else if f <= 0 {
			fmt.Println("нет 0 и отрицателных чисел")
		} else if f <= 10 {
			fmt.Println(roman[f-1])
		} else if f == 20 {
			fmt.Println("XX")
		}
	}
	if !isRoman {
		fmt.Println(f)
	}
}
