package main

import "fmt"

func main() {

	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}

	for i, num := range pontosPlanningPoker {
		fmt.Println("posicao", i, "numero", num)
	}

}
