package main

import (
	rect "demo/shape/rect" // nameof the module/root package/sub package
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

}
