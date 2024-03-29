package main

import "fmt"

func getGun(name string) (IGun, error) {
	if name == "AK47" {
		return newAK47(), nil
	}
	if name == "Musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("unknown name %s passed", name)
}