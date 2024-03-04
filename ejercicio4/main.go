package main

import (
	"errors"
	"fmt"
)

/*
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no
alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro).
*/
type CumstomErrorSalary struct {
	msg string
}

func (e CumstomErrorSalary) Error() string {
	return e.msg
}

func main() {
	salary := 9000
	err := salaryCustomError(salary)
	var errsalary *CumstomErrorSalary // errMath -> 0x000a | nil

	// ChangeSomething(&errMath)

	if errors.As(err, &errsalary) { // errors.As(err, errMath): e -> 0x000b | nil
		switch errsalary.msg {
		case "Error: the minimum taxable amount is 150,000 and the salary entered is:":
			fmt.Println("Comapro los errores")
		default:
			fmt.Println("No se compararon")
		}
	}
}

func salaryCustomError(salary int) (err error) {
	if salary <= 10000 {
		err = &CumstomErrorSalary{msg: "Error: the minimum taxable amount is 150,000 and the salary entered is:"}
		err = fmt.Errorf("%w. %w", err, "Indeterminate division")
	}
	return
}
