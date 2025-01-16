package main

import "fmt"

func main() {
	countTestCase := 0
	fmt.Scan(&countTestCase)

	for i := 0; i < countTestCase; i++ {
		var str string
		fmt.Scan(&str)

		if !exception(str) {
			fmt.Println("NO")
		} else {
			worker(str)
		}
	}
}

func exception(str string) bool {
	if len(str) == 0 || str[0] != 'M' {
		return false
	}

	lenStr := len(str)

	for i := 0; i < lenStr-1; i++ {
		if str[i] == str[i+1] {
			return false
		}
	}

	return true
}

func worker(str string) {
	mPointer := 0
	flag := false
	lenStr := len(str)

	for i := 0; i < lenStr; i++ {
		if str[i] == 'M' {
			mPointer++
		} else if str[i] == 'R' {
			if i+1 < lenStr && str[i+1] == 'C' {
				mPointer--
				i++
			} else {
				flag = true
				break
			}
		} else if str[i] == 'C' {
			if i+1 < lenStr && str[i+1] == 'M' {
				i++
			} else {
				flag = true
				break
			}
		} else {
			mPointer--
		}

		if mPointer < 0 {
			fmt.Println("NO")
			return
		}
	}

	if mPointer == 0 && !flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
