package factory

import (
	"github.com/Design-Patterns/abstractfactory/shoe"
	"github.com/Design-Patterns/abstractfactory/shirt"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() shoe.IShoe {
	return &shoe.AdidasShoe{
		Shoe: shoe.Shoe{
			Name: "Adidas Shoe",
			Price: 2000.0,
		},
	}
}

func (a *Adidas) MakeShirt() shirt.IShirt {
	return &shirt.AdidasShirt{
		Shirt: shirt.Shirt{
			Size: 40,
			Price: 1500.0,
		},
	}
}