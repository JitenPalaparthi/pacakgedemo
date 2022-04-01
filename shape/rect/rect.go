package rect

import "errors"

type Rect struct {
	Length, Bredth, Area, Perimeter float32
}

// This is a function
func AreaOfRect(r Rect) float32 {
	return r.Bredth * r.Length
}

// Methods on types.
// A method on a type must have a receiver of that type
// If a receiver is a pointer then it will call the same rect object
// else if it is a value type then it works on a new copy
func (r *Rect) GetArea() (float32, error) {
	if r.Length <= 0 || r.Bredth <= 0 {
		return 0, errors.New("invalid Length and Bredth values")
	}
	r.Area = r.Length * r.Bredth
	return r.Area, nil

}
func (r *Rect) GetPerimeter() (float32, error) {
	if r.Length <= 0 || r.Bredth <= 0 {
		return 0, errors.New("invalid Length and Bredth values")
	}
	r.Perimeter = 2 * (r.Length + r.Bredth)
	return r.Perimeter, nil
}
