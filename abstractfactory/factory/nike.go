package factory

import (
	"github.com/Design-Patterns/abstractfactory/shoe"
	"github.com/Design-Patterns/abstractfactory/shirt"
)

type Nike struct {
}

func (a *Nike) MakeShoe() shoe.IShoe {
	return &shoe.NikeShoe{
		Shoe: shoe.Shoe{
			Name: "Nike Shoe",
			Price: 1800.0,
		},
	}
}

func (a *Nike) MakeShirt() shirt.IShirt {
	return &shirt.NikeShirt{
		Shirt: shirt.Shirt{
			Size: 40,
			Price: 1600.0,
		},
	}
}