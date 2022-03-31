package main

import (
	"demo/area"           //nameof the module/root package
	rect "demo/area/rect" // nameof the module/root package/sub package
	"fmt"                 // a standard pacakge
)

func main() {
	fmt.Println("hello package")
	a := rect.AreaOfRectangle(10.23, 15.34)
	//area := area.areaOfRect(10.23, 15.34)
	//area.L = 100.12
	//area.B = 101.32
	fmt.Println(a)
	fmt.Println(rect.L, rect.B)
	areaS := area.AreaOfSquare(12.34)
	fmt.Println(areaS)
}
