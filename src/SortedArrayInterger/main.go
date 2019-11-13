package main

func main() {
	perOne := Permanent{
		empId:    123,
		basicpay: 100,
		pf:       500,
	}

	contract := Contract{
		empId:    123,
		basicpay: 200,
	}

	var slice  = []SalaryCalculator{perOne, contract}
	totalExpense(slice)
}
