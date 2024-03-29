package main

import "fmt"

func main() {
	ak47, _ := getGun("AK47")
	musket, _ := getGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Println("name: %s", g.getName())
	fmt.Println("power: %d", g.getPower())
}