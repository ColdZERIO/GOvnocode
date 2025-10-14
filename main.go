package main

import (
	hui "GOvnotest/Hui"
	zlp "GOvnotest/Zlp"
	"fmt"
)

func main() {
	num1 := zlp.Mul(5, 8)

	num2 := zlp.Div(30, 0) // Дим, чет тут сломано, чекни файл в пакете Zlp

	fmt.Println(num1, num2)

	zlp.Out1()

	Zslice := []int{5, 3, 7, 8, 4, 5, 75, 34, 7, 5, 543, 4, 6, 5, 7, 4, 45, 6}

	// Сань, какая-то залупа с этой функцией. Чекни папку Hui
	res := hui.ChekSLice(Zslice)
	fmt.Println(res)

	hui.ChekSLice() // Сань, какая-то залупа с этой функцией. Чекни папку Hui

	err := zlp.Div()

	if err != nil {
		fmt.Println(err)
	}
}
