package main

import "fmt"

var (
	array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = []struct {
		n int
		s string
	}{
		{0, "odd"},
		{1, "even"},
		{2, "odd"},
		{3, "even"},
		{4, "odd"},
	}

	slice  = array[5:8]
	start  = array[:6] //same cap
	end    = array[7:] // change cap
	all    = array[:]
	extend = start[3:]
)

func main() {
	array[7] = 777
	fmt.Println(array)
	fmt.Println(slice)

	printSlice("Start = ", start)
	printSlice("End = ", end)
	printSlice("All = ", all)
	printSlice("Extend = ", extend)

	a := make([]int, 5)
	b := make([]int, 0, 5) //type, len, cap

	c := b[:2]
	d := c[2:5]

	printSliceMake("a = ", a)
	printSliceMake("b = ", b)
	printSliceMake("c = ", c)
	printSliceMake("d = ", d)

}

func printSlice(sliceName string, v []int) {
	fmt.Println(sliceName, "len = ", len(v), "cap = ", cap(v))
}

func printSliceMake(sliceName string, v []int) {
	fmt.Println(sliceName, v, "len = ", len(v), "cap = ", cap(v))
}
