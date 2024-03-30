package factory

import (
	"fmt"

	"github.com/Design-Patterns/abstractfactory/shirt"
	"github.com/Design-Patterns/abstractfactory/shoe"
)

type ISportsFactory interface {
	MakeShoe() shoe.IShoe
	MakeShirt() shirt.IShirt
}

func GetSportsFactory(factoryName string) (ISportsFactory, error) {
	if factoryName == "Adidas" {
		return &Adidas{}, nil
	}
	if factoryName == "Nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("unknown factory name: %s", factoryName)
}