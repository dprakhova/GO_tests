package main

import "fmt"

func calculateTax(salary float64) (float64, float64) {
	var tax, remainder float64
	if salary <= 1000 {
		tax = salary * 0.05
	} else if salary <= 3000 {
		tax = 1000*0.05 + (salary-1000)*0.08
	} else if salary <= 5000 {
		tax = 1000*0.05 + 2000*0.08 + (salary-3000)*0.1
	} else {
		tax = 1000*0.05 + 2000*0.08 + 2000*0.1 + (salary-5000)*0.2
	}
	remainder = salary - tax
	return tax, remainder
}

func main() {
	salary := 1500.0
	tax, remainder := calculateTax(salary)
	fmt.Printf("Salary: %.2f\nTax: %.2f\nRemainder: %.2f\n", salary, tax, remainder)
}
