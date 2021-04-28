package entities

import (
	"errors"
	"strconv"
	"testing"
)

type CreateItemResult struct {
	name          string
	minLevel      int32
	maxLevel      int32
	consumable    bool
	expectedError error
}

var itemCreationResults = []CreateItemResult{
	{"TestItem", 1, 10, true, nil},
	{"", 1, 10, true, &InvalidNameError{}},
	{"TestItem", 0, 15, true, &InvalidLevelError{}},
	{"TestItem", 1, 21, true, &InvalidLevelError{}},
}

func TestCreateItem(t *testing.T) {
	for _, currentTest := range itemCreationResults {
		createdItem, err := CreateItem(currentTest.name, currentTest.minLevel, currentTest.maxLevel, currentTest.consumable)

		if currentTest.expectedError != nil {
			// This test should have returned us an error
			if errors.Is(err, currentTest.expectedError) {
				t.Fatalf("Expected error: %s - Received Error: %s", currentTest.expectedError, err)
			}
		} else {
			// This test should have succcessfully create the item
			if currentTest.name != createdItem.ItemName {
				t.Fatalf("Expected itemName: %s - Received itemName: %s", currentTest.name, createdItem.ItemName)
			}

			if currentTest.minLevel != createdItem.MinLevel {
				t.Fatalf("Expected itemMinLevel: %d - Received itemMinLevel: %d", currentTest.minLevel, createdItem.MinLevel)
			}

			if currentTest.maxLevel != createdItem.MaxLevel {
				t.Fatalf("Expected itemMaxLevel: %d - Received itemMaxLevel: %d", currentTest.maxLevel, createdItem.MaxLevel)
			}

			if currentTest.consumable != createdItem.Consumable {
				t.Fatalf("Expected itemConsumable: %s - Received itemConsumable: %s", strconv.FormatBool(currentTest.consumable), strconv.FormatBool(createdItem.Consumable))
			}
		}
	}

}
