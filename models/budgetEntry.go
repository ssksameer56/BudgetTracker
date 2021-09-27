package models

import "errors"

type BudgetEntry struct {
	Name        string
	Description string
	Amount      int
	Category    string
}

func (b *BudgetEntry) UpdateAmount(value int) error {
	if value == 0 {
		return errors.New("Cannot have zero amount")
	}
	b.Amount = value
	return nil
}

func (b *BudgetEntry) UpdateName(value string) error {
	b.Name = value
	return nil
}

func (b *BudgetEntry) UpdateDesc(value string) error {
	b.Description = value
	return nil
}

func (b *BudgetEntry) UpdateCategory(value string) error {
	if ContainsCategory(value) {
		b.Category = value
		return nil
	} else {
		return errors.New("Category Not Valid")
	}
}

func GetNewBudgetEntry(amt int, name, desc, cat string) (BudgetEntry, error) {
	if amt == 0 {
		return BudgetEntry{}, errors.New("Amount is zero")
	}
	entry := BudgetEntry{
		Name:        name,
		Description: desc,
		Category:    cat,
		Amount:      amt,
	}
	return entry, nil
}
