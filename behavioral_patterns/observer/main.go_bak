package main

func main() {

	shirtItem := newItem("Nike Shirt")
	shopesItem := newItem("adidas shoes")
	coatItem := newItem("adidas Coat")

	observerFirst := &customer{id: "123@gmail.com"}
	observerSecond := &customer{id: "second@gmail.com"}
	observerTom := &customer{id: "tom@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.register(observerTom)

	coatItem.register(observerSecond)

	//shopesItem.register(observerFirst)
	//shopesItem.register(observerSecond)
	shopesItem.register(observerTom)

	shirtItem.updateAvailability()
	coatItem.updateAvailability()
	shopesItem.updateAvailability()

}
