package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId  int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {

	for _, v := range s {
		expense := 0
		expense = expense + v.CalculateSalary()
		fmt.Println()
		fmt.Printf("Total Expense Per Month $%d", expense)
	}

}