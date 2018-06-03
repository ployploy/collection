package main

import (
	"fmt"
)

func main() {
	mymap := make(map[int]int)
	mymap[1] = 1
	mymap[2] = 2

	fmt.Println(mymap[3])
	if mymap[3] != 0 {
		fmt.Println(mymap[3])
	}
	// check ให้เห็นว่าจริงๆ  2 มันมีค่า defualt เลยเป็น true
	// value, ok := mymap[2]
	// fmt.Println(value, ok)
	//ok ค่าตรงนี้มันเป็น defualt = 0 เพราะ 3 ไม่ได้ถูกกำหนดไว้
	if value, ok := mymap[3]; ok {
		fmt.Println(value)
	}
}
