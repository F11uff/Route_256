package main

import (
	"fmt"
)

func main() {
	salary := "0"
	countTest := 0
	fmt.Scan(&countTest)
	for i := 0; i < countTest; i++ {
		fmt.Scan(&salary)
		delete(salary)
	}

}

func delete(salary string) {
	if ok := exceptions(salary); !ok {
		findMin(salary)
	}
}

func exceptions(salary string) bool {
	if len(salary) < 2 {
		fmt.Println(0)
		return true
	}

	return false
}

func findMin(salary string) {
	lenString := len(salary)
	i := 0

	for ; i < lenString-1; i++ {
		if salary[i] >= salary[i+1] {
			fmt.Print(string(salary[i]))
			continue
		} else {
			break
		}
	}

	i++

	for ; i < lenString; i++ {
		fmt.Print(string(salary[i]))
	}

	fmt.Println()
}
