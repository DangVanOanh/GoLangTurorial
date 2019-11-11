package main

import (
	"SortedArray"
	"fmt"
)

func main()  {
	var  sliceSort  = []int{-1,2,4,6,3,40,20,-5}
	fmt.Println(sliceSort)
	SortedArray.BubbleSortArray(sliceSort)
	fmt.Println(sliceSort)
}


