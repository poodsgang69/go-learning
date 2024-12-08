package main

import (
	"fmt"
)

func main() {
	// solve()
	var x int = 5
	summ(x)
	var arr [5]int = [5]int{100, 200, 300, 400, 500}
	fmt.Println(rev(arr))
	var c string = "Pinnu"
	fmt.Println(switchT(c))
	customBreak()
	var slc []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deferT(slc))
}

func solve() {
	fmt.Println("Yo")
}

func summ(x int) {
	for i := 0; i < 10; i++ {
		x += x
	}
	arr := []int{10, 20, 30, 40}
	for idx, val := range arr {
		fmt.Printf("Value: %d occurs at index: %d\n", val, idx)
	}
	// fmt.Println(x)
}

func rev(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func switchT(c string) bool {
	switch c {
	case "chinnu", "Chinnu", "CHINNU":
		return true
	}
	return false
}

func customBreak() {
Loop:
	for i := 0; i < 10; i++ {
		switch {
		case i < 5:
			fmt.Print("Less than 5")
		case i > 5:
			fmt.Print("Greater than 5")
		case i == 5:
			break Loop
		}

	}
}

func deferT(slc []int) (text string) {
	defer show("The Method has exit.")

	if len(slc)%2 == 0 {
		return "Slice is Even"
	} else {
		return "Slice is Odd"
	}

}

func show(text string) {
	fmt.Println(text)
}
