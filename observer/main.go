package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observer1 := &Customer{id: "customer1@abc.com"}
	observer2 := &Customer{id: "customer2@abc.com"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()

}
