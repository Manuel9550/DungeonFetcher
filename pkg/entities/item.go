package entities

import "github.com/rs/xid"

type Item struct {
	ID         string `json:"id"`
	ItemName   string `json:"name"`
	MinLevel   int32  `json:"minLevel,omitempty"`
	MaxLevel   int32  `json:"maxLevel,omitempty"`
	Consumable bool   `json:"consumable,omitempty"`
}

type InvalidLevelError struct{}

func (e *InvalidLevelError) Error() string {
	return "item levels must be between 1 and 20"
}

type InvalidNameError struct{}

func (e *InvalidNameError) Error() string {
	return "items must have a non blank name"
}

func CreateItem(name string, minLevel int32, maxLevel int32, consumable bool) (Item, error) {

	// level checks
	if minLevel <= 0 || minLevel > 20 {
		return Item{}, &InvalidLevelError{}
	}

	if maxLevel <= 0 || maxLevel > 20 {
		return Item{}, &InvalidLevelError{}
	}

	// blank name check
	if name == "" {
		return Item{}, &InvalidNameError{}
	}

	return Item{
		ID:         xid.New().String(),
		ItemName:   name,
		MinLevel:   minLevel,
		MaxLevel:   maxLevel,
		Consumable: consumable,
	}, nil
}
