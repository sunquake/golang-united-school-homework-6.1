package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity > len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("full")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i <= b.shapesCapacity {
		return b.shapes[i], nil
	}
	return Shape{}, fmt.Errorf("out of bounds")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var s Shape
	s, err := b.GetByIndex(i)
	if  err != nil {
		return s, err
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var s Shape
	s, err := b.GetByIndex(i)
	if  err != nil {
		return s, err
	}
	b.shapes[i] = shape
	return s, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	s := 0.
	for _, v := range b.shapes {
		s += v.CalcPerimeter()
	}
	return s
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	s := 0.
	for _, v := range b.shapes {
		s += v.CalcArea()
	}
	return s
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	c := 0
	s := Shape{}
	for i, v := range b.shapes {
		if s, ok := b.(Circle); ok {
			copy(b.shapes[i:], b.shapes[i+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			c++
		}
	}
	if c != 0 {return nil}
	return fmt.Errorf("no circles")
}
