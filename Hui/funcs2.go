package hui // это не то что кажется.

var count int

func ChekSLice(slice []int) int { // Нужна функция которая ищет количество цифр "5" в слайсе

	for _, val := range slice {
		if val == 5 {
			count++
		}
	}
	return count
}

// Ах да, и он не видит почему-то эту функцию в мейне... исправь этоу залупу!
