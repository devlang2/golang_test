package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// array
	spew.Println("=================== array")
	intArr := [3]int{1, 2, 3}
	changeIntArr(intArr)
	spew.Dump(intArr)

	// slice **
	spew.Println("=================== slice")
	sliceIntArr := make([]int, 0, 10)
	sliceIntArr = append(sliceIntArr, 1, 2, 3)
	changeSliceIntArr(sliceIntArr)
	spew.Dump(sliceIntArr)

	// struct
	spew.Println("=================== struct")
	p := Person{"a", 1}
	changeStruct(p)
	spew.Dump(p)

	// struct array slice **
	spew.Println("=================== struct array slice")
	p1 := Person{"a", 1}
	p2 := Person{"b", 2}
	pArrSlice := []Person{p1, p2}
	changeStructArrSlice(pArrSlice)
	spew.Dump(pArrSlice)

	// struct array
	spew.Println("=================== struct array")
	pArr := [2]Person{p1, p2}
	changeStructArr(pArr)
	spew.Dump(pArr)
	changeStructArr(pArr)
	spew.Dump(pArr)

	// struct array
	spew.Println("=================== map")
	m := map[string]string{
		"name": "won",
		"addr": "korea",
	}
	changeMap(m)
	spew.Dump(m)

}

func changeMap(m map[string]string) {
	m["name"] = "xxxxxxxxxxxxxx"
}

func changeIntArr(arr [3]int) {
	arr[0] = 999
}

func changeSliceIntArr(arr []int) {
	arr[0] = 999
}

func changeStruct(p Person) {
	p.Age = 999
}

func changeStructArrSlice(arr []Person) {
	arr[0].Age = 999
}
func changeStructArr(arr [2]Person) {
	arr[0].Age = 999
}
