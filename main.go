package main
// Залупа в ветке ZalupnieHui
// 2
import (
	hui "GOvnotest/Hui"
	zlp "GOvnotest/Zlp"
	"fmt"
)

func main() {
	num1 := zlp.Mul(5, 8)

	num2, err := zlp.Div(30, 0) // Дим, чет тут сломано, чекни файл в пакете Zlp

	fmt.Println(num1, num2)

	zlp.Out1()

	Zslice := []int{5, 3, 7, 8, 4, 5, 75, 34, 7, 5, 543, 4, 6, 5, 7, 4, 45, 6}

	// Сань, какая-то залупа с этой функцией. Чекни папку Hui
	res := hui.ChekSLice(Zslice)
	fmt.Println(res)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num2)
	}

	for i := range 100 {
		fmt.Println(i)
	}
}
