package main

import "fmt"

var (
	array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = []struct {
		n int
		s string
	}{
		{n: 0, s: "odd"},
		{n: 1, s: "even"},
		{n: 2, s: "odd"},
		{n: 3, s: "even"},
		{n: 4, s: "odd"},
	}

	slice  []int = array[5:8]
	start  []int = array[:6] //same cap
	end    []int = array[7:] // change cap
	all    []int = array[:]
	extend []int = start[3:]

	a = make([]int, 5)
	b = make([]int, 0, 5) //type, len, cap

	c []int = b[:2]
	d []int = c[2:5]
)

func main() {
	array[7] = 777
	fmt.Println(array)
	fmt.Println(slice)

	printSlice("Start =", start)
	printSlice("End =", end)
	printSlice("All =", all)
	printSlice("Extend =", extend)

	printSliceMake("a =", a)
	printSliceMake("b =", b)
	printSliceMake("c =", c)
	printSliceMake("d =", d)

}

func printSlice(sliceName string, v []int) {
	fmt.Println(sliceName, v, "len = ", len(v), "cap = ", cap(v))
}

func printSliceMake(sliceName string, v []int) {
	fmt.Println(sliceName, v, "len = ", len(v), "cap = ", cap(v))
}
