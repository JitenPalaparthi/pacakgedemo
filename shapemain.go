// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func main() {

	e := &Empty{}
	r := &Rect{L: 12.34, B: 12.78}
	s := &Square{S: 25.25}
	c := &Cube{L: 12.45, B: 12.78, H: 12.1}
	fmt.Println("Enter 0.Empty 1.Rect 2.Square 3.Cube")
	var ch uint8
	fmt.Scan(&ch)
	switch ch {
	case 0:
		Area(e)
	case 1:
		Area(r)
	case 2:
		Area(s)
	case 3:
		Area(c)
	default:
		fmt.Println("Wrong input")
	}

	iarray := make([]IShape, 4)
	iarray[0] = e
	iarray[1] = r
	iarray[2] = s
	iarray[3] = c

	for _, iface := range iarray {
		fmt.Println(iface.Area())
	}

	objs := make([]interface{}, 5)
	objs[0] = e
	objs[1] = r
	objs[2] = s
	objs[3] = c
	objs[4] = &NoIShape{}
	for _, obj := range objs {
		switch obj.(type) {
		case IShape:
			fmt.Println(obj.(IShape).Area())

		default:
			fmt.Println("Not implemented stuff", reflect.TypeOf(obj))
		}
	}

}

func Area(ishape IShape) {
	area, err := ishape.Area()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area:", area)
	}
}

type Empty struct{} // empty struct

type NoIShape struct{}

func (ni *NoIShape) SayHai() {
	fmt.Println("Hello..")
}

func (r *Empty) Area() (float32, error) {
	return 0, nil
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() (float32, error) {
	if r == nil || r.L <= 0 || r.B <= 0 {
		return 0, fmt.Errorf("Wrong L and B values")
	}
	return r.L * r.B, nil
}

type Square struct {
	S float32
}

func (s *Square) Area() (float32, error) {
	if s == nil || s.S <= 0 {
		return 0, fmt.Errorf("Wrong S value")
	}
	return s.S * s.S, nil
}

type Cube struct {
	L, B, H float32
}

func (r *Cube) Area() (float32, error) {
	if r == nil || r.L <= 0 || r.B <= 0 || r.H <= 0 {
		return 0, fmt.Errorf("Wrong L or B or H values")
	}
	return r.L * r.B * r.H, nil
}

type IShape interface {
	Area() (float32, error)
	//Perimeter() (float32, error)
}
