package main

import "fmt"

type Car interface {
	start()
}

type Tesla struct {
}

func (r Tesla) start() {
	fmt.Println("Starting your Tesla!")
}

type BMW struct {
}

func (r BMW) start() {
	fmt.Println("Starting your BMW")
}

type Button struct {
	car Car
}

func (r Button) press() {
	r.car.start()
}

func main() {
	tesla := Tesla{}
	bmw := BMW{}

	button := Button{
		car: tesla,
	}
	button.car.start()
	button.car = bmw
	button.car.start()
}
