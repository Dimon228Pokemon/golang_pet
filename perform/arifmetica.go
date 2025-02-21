package perform

import (
	"errors"
	"fmt"
)

type Do interface {
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
}

type Number struct {
	Value   float64
	Is_zero bool
}

func Thinking() {
	fmt.Println("...")
}

func Addition(first_number Number, second_number Number) float64 {
	result := first_number.Value + second_number.Value
	return result
}

func Subtraction(first_number Number, second_number Number) float64 {
	result := first_number.Value - second_number.Value
	return result
}

func Multiplication(first_number Number, second_number Number) float64 {
	result := first_number.Value * second_number.Value
	return result
}

func (n Number) Check_zero() bool {
	return n.Is_zero
}

func Division(first_number Number, second_number Number) (float64, error) {
	if second_number.Check_zero() {
		return 0, errors.New("Нельзя делить на ноль")
	}
	result := first_number.Value / second_number.Value
	return result, nil
}
