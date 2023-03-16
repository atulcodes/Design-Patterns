package main

type Observer interface {
	update(string)
	getId() string
}
