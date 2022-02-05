package main

import (
	"fmt"
	"interfaces/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("Robert")

	total := frank.CalcTotal(90, 65, 93.6)

	fmt.Printf("\nThe employee %v \n \n and the total is %v \n\n", frank, total)

	robert.DeactivateEmployee(frank)

	total2 := frank.CalcTotal(90, 65, 93.6)

	fmt.Printf("\nThe total is %v", total2)

	robert.ReactivateEmployee(frank)

	total3 := frank.CalcTotal(100, 65, 93.6)

	fmt.Printf("\nThe total is %v", total3)
}
