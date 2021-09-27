package main

import (
	"BudgetTracker/services"
	"fmt"
)

func main() {

	menuOptions := "Enter 1 for adding Entry, 2 for viewing all entries, 3 for seeing all the totals, 4 to exit"
	bService := services.GetNewBudgetService()
	var op int
	for {
		fmt.Println(menuOptions)
		_, err := fmt.Scanln(&op)
		if err != nil {
			continue
		}
		switch op {
		case 1:
			fmt.Println("Enter Name, amount, Description, Category(Credit/Debit)")
			var amt int
			var name, desc, cat string
			fmt.Scanln(&name, &amt, &desc, &cat)
			err := bService.AddEntry(amt, name, desc, cat)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 2:
			allEntries := bService.GetAllEntries()
			for _, v := range allEntries {
				fmt.Println(v.Name, v.Amount, v.Category, v.Description)
			}
		case 3:
			fmt.Println("Enter Category You Want")
			var cat string
			fmt.Scanln(&cat)
			val, err := bService.TotalByCategory(cat)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(val)
		case 4:
			break
		}
	}
}
