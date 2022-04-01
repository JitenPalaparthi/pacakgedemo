package main

import (
	rect "demo/shape/rect" // nameof the module/root package/sub package
	"demo/types"
	"fmt"
)

func main() {
	r := &rect.Rect{Length: 12.45, Bredth: 13.65}
	area, err := r.GetArea()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of Rect:", area)
	}
	perimeter, err := r.GetPerimeter()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of Perimeter:", perimeter)
	}

	fmt.Println(r.Area)
	fmt.Println(r.Perimeter)

	fmt.Println("Calling using rect package function", rect.AreaOfRect(*r))
	// using methods on userdefined types

	var mi types.MyInt

	mi = 100
	var str string = mi.ToString()
	fmt.Println(str)

	var i float32 = 12312.312
	//var i int = 100
	str1 := types.MyInt(i).ToString()
	fmt.Println(str1)
}
