package main

import "fmt"

func main() {

	var strings [4]string
	strings[0] = "Kotko"
	strings[1] = "Graniaste"
	strings[2] = "Cztero"
	strings[3] = "Kanciaste"
	for i, word := range strings {
		fmt.Println(i, word)
	}

	numbers := [4]int{10, 20, 30, 40}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// no need to enter the number of items
	// if we initialize with values
	b := [...]string{"Penn", "Teller"}

	fmt.Println(b)

}
