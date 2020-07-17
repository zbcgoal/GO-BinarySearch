package main

import (
	"fmt"
	"os"
	"sort"
)

func BinarySearch(ary *[]int, leftindex int, rightindex int, valueFind int)  {
	mid := (leftindex + rightindex)/2
	if leftindex > rightindex{
		fmt.Println("can not find")
		os.Exit(1)
	}
	if (*ary)[mid] > valueFind {
		BinarySearch(ary,leftindex,mid-1,valueFind)
	}else if (*ary)[mid] < valueFind {
		BinarySearch(ary,mid+1,rightindex,valueFind)
	}else {
		fmt.Printf("要找的数为%d,找到了，其下标为%d",valueFind,mid)
	}
}

func main()  {
	var ary []int
	ary = []int{100,22,33,45,32,67,42,78,31,9,1,4,43}
	//从小到大排序
	sort.Ints(ary)
	fmt.Println(ary)
	fmt.Println(len(ary))
	BinarySearch(&ary, 0, len(ary) - 1, 100 )
}



