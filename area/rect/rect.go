package rect

import "fmt"

var (
	L float32
	B float32
)

func init() {
	L = 10
	B = 10
	fmt.Println("init is a default function. is called automatially")
}

func init() {
	//L = 0.0
	//B = 0.0
	fmt.Println("init is a default function. is called automatially")
}

func areaOfRect(l, b float32) float32 {
	return l * b
}

func AreaOfRectangle(l, b float32) float32 {
	return areaOfRect(l, b)
}
