package zlp

import "errors"

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) { // Я СДЕЛЯЛЬ
	if b == 0 {
		return 0, errors.New("division by zero DOLBAYOB")
	}
	return a / b, nil
}
