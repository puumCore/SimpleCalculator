package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var x, y float64
	var sign string

	fmt.Println("Enter the first number: ")
	fmt.Scan(&x)
	fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scan(&sign)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&y)

	cal := Draft[float64, float64]{
		x, y, sign,
	}

	result, err := calculate(cal)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("## %.2f %s %.2f = %.2f\n", x, sign, y, result)
	}

}

type Draft[F Number, S Number] struct {
	firstNumb  F
	secondNumb S
	operator   string
}

type Number interface {
	int | int64 | float32 | float64 | complex64 | complex128
}

func calculate(d Draft[float64, float64]) (float64, error) {
	switch d.operator {
	case "+":
		return d.firstNumb + d.secondNumb, nil
	case "-":
		return d.firstNumb - d.secondNumb, nil
	case "*":
		return d.firstNumb * d.secondNumb, nil
	case "/":
		if d.secondNumb == 0 {
			return 0, fmt.Errorf("you can't divide by 0")
		}
		return d.firstNumb / d.secondNumb, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}
