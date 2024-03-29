package main

type IGun interface {
	setName(string)
	getName() string
	setPower(int) 
	getPower() int 
}