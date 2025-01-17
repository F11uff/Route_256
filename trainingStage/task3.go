package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	countTestCaseStr, _ := reader.ReadString('\n')
	countTestCase, _ := strconv.Atoi(strings.TrimSpace(countTestCaseStr))
	flag := true

	for i := 0; i < countTestCase; i++ {
		sizeArrayStr, _ := reader.ReadString('\n')
		sizeArray, _ := strconv.Atoi(strings.TrimSpace(sizeArrayStr))

		strArray, _ := reader.ReadString('\n')
		strArray = strings.TrimSpace(strArray)

		input, _ := reader.ReadString('\n')

		if input[0] == ' ' {
			flag = false
		}

		input = strings.TrimSpace(input)

		if !flag {
			fmt.Println("no")
			continue
		}

		validator(sizeArray, strArray, input)
	}
}

func validator(sizeArray int, strArray, input string) {
	if len(strArray) != len(input) {
		fmt.Println("no")
		return
	}

	sliceStr := split(strArray)

	sort.Slice(sliceStr, func(i, j int) bool {
		numI, _ := strconv.Atoi(sliceStr[i])
		numJ, _ := strconv.Atoi(sliceStr[j])
		return numI < numJ
	})

	sliceInput := strings.Split(input, " ")

	if sizeArray != len(sliceInput) {
		fmt.Println("no")
		return
	}

	for i := 0; i < len(sliceStr); i++ {
		if sliceStr[i] != sliceInput[i] {
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
}

func split(strArray string) []string {
	return strings.Split(strArray, " ")
}
