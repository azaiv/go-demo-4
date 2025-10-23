package main

import (
	"fmt"
)

type intArray = [3]int

func main() {
	a := intArray{1, 2, 3}
	//pointerA := &a // Адрес в памяти
	//multiple(&a)
	//fmt.Println(*pointerA) // Получение занчение из адреса
	fmt.Println(a)
	reversed(&a)
	fmt.Println(a)
}

func reversed(arr *intArray) {
	for index, value := range *arr {
		arr[len(arr)-1-index] = value
	}
}
