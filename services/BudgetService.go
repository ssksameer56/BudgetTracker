package services

import (
	"BudgetTracker/models"
	"errors"
	"strings"
)

type Entry = models.BudgetEntry

type IBudgetService interface {
	AddEntry(int, string, string, string)
	UpdateEntry(*Entry, string, string, string, int)
	UpdateAmount(*Entry, int)
	UpdateDesc(*Entry, string)
	UpdateName(*Entry, string)
}

type BudgetService struct {
	entries []*Entry
}

var entries = []*Entry{}

func GetNewBudgetService() *BudgetService {
	newService := BudgetService{[]*Entry{}}
	return &newService
}

func (*BudgetService) AddEntry(amt int, name, desc, cat string) error {
	newEnt, err := models.GetNewBudgetEntry(amt, name, desc, cat)
	if err != nil {
		return err
	}
	entries = append(entries, &newEnt)
	return nil
}

func (*BudgetService) UpdateAmount(entry *Entry, amt int) error {
	err := entry.UpdateAmount(amt)
	if err != nil {
		return err
	}
	return nil
}

func (*BudgetService) UpdateDesc(entry *Entry, desc string) error {
	err := entry.UpdateDesc(desc)
	if err != nil {
		return err
	}
	return nil
}

func (*BudgetService) UpdateCategory(entry *Entry, cat string) error {
	err := entry.UpdateCategory(cat)
	if err != nil {
		return err
	}
	return nil
}

func (*BudgetService) TotalByCategory(cat string) (int, error) {
	sum := 0
	if !models.ContainsCategory(cat) {
		return 0, errors.New("Invalid category")
	}
	for _, v := range entries {
		if strings.Compare(cat, v.Category) == 0 {
			sum += v.Amount
		}
	}
	return sum, nil
}

func (*BudgetService) GetAllEntries() []*Entry {
	return entries
}
