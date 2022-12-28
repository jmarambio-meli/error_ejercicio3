package main

import (
	"errors"
	"fmt"
)

var (
	Error = errors.New("Error: El salario es menor a 10000")
)

func main() {
	var salary int = 15000
	err := validarSalario(salary)
	if err != nil {
		if errors.Is(err, Error) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar Impuesto")
	}
	fmt.Println(errors.Is(err, Error))

}

func validarSalario(salary int) error {
	if salary < 10000 {
		return Error
	}

	return nil
}
