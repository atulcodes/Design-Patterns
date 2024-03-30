package main

import (
	"fmt"

	"github.com/Design-Patterns/abstractfactory/factory"
	"github.com/Design-Patterns/abstractfactory/shirt"
	"github.com/Design-Patterns/abstractfactory/shoe"
)

func main() {
	nikeFactory, _ := factory.GetSportsFactory("Nike")
	adidasFactory, _ := factory.GetSportsFactory("Adidas")
	
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	aididasShirt := adidasFactory.MakeShirt()

	printShirtDetails(nikeShirt)
	printShirtDetails(aididasShirt)

	printShoeDetails(nikeShoe)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(s shoe.IShoe) {
	fmt.Printf("Shoe name: %s\n", s.GetName())
	fmt.Printf("Shoe price: %f\n", s.GetPrice())
}

func printShirtDetails(s shirt.IShirt) {
	fmt.Printf("Shirt size: %d\n", s.GetSize())
	fmt.Printf("Shirt price: %f\n", s.GetPrice())
}