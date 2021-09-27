package models

import (
	"errors"
	"strings"
)

var categories = []string{"Credit", "Debit"}

func GetAllCategories() []string {
	return categories
}

func ContainsCategory(val string) bool {
	val = strings.ToLower(val)
	for _, v := range categories {
		lowerV := strings.ToLower(v)
		if strings.Compare(val, lowerV) == 0 {
			return true
		}
	}
	return false
}

func AddCategory(val string) error {
	if len(val) == 0 {
		return errors.New("empty string not allowed")
	}
	categories = append(categories, val)
	return nil
}
