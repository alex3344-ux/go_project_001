package main


import (
	"fmt"
)


/*

func main() {
	array := [3]string{"1", "2", "3"} // Это массив
	slice := []string{"4", "5", "6"} // Это слайс

	fmt.Println("Это Массив - ", array)
	fmt.Println("Это Слайс - ", slice)
}

*/

func main() {
	matrix := make([][]int, 10)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x
		}
		fmt.Println(matrix[x])
	}

}






















