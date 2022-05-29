package golang_united_school_homework

import (
	"errors"
)

var errIndex = errors.New("index doesn't exist or index went out of the range")

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

	if len(b.shapes)+1 > b.shapesCapacity {
		return errors.New("shape's length goes out of shapeCapacity range")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	for index, value := range b.shapes {
		if index == i {
			return value, nil
		}
	}
	return nil, errIndex

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	for indx, val := range b.shapes {
		if indx == i {
			b.shapes = append(b.shapes[:indx], b.shapes[i+1:]...)

			return val, nil
		}
	}

	return nil, errIndex

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	for inx := range b.shapes {
		if inx == i {
			val := b.shapes[i]
			b.shapes[i] = shape
			return val, nil
		}
	}

	return nil, errIndex

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, val := range b.shapes {
		sum += val.CalcPerimeter()

	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, val := range b.shapes {
		sum += val.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	for _, v := range b.shapes {
		if _, ok := v.(*Circle); ok {
			exist = true
			break
		}
	}

	if exist {
		indexes := []int{}
		for i, v := range b.shapes {
			if _, ok := v.(*Circle); !ok {
				indexes = append(indexes, i)
			}
		}

		temp := []Shape{}
		for _, v := range indexes {
			temp = append(temp, b.shapes[v])
		}
		b.shapes = temp

		return nil
	} else {

		return errors.New("circles are not exist in the list")
	}

}
