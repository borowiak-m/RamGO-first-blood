package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println
var pf = fmt.Printf
var typeOf = reflect.TypeOf

type Aniamal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat attacks")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says hissssss")
}

func (c Cat) HappySound() {
	pl("Cat says purrrrrrrr")
}

type customer struct {
	name    string
	address string
	bal     float64
}

type rectangle struct {
	length, height float64
}

func getCustomerInfo(c customer) {
	pf("%s owes us %.2f\n", c.name, c.bal)
	pf("Customer lives at : %s\n", c.address)
}

func updateCustAddress(c *customer, address string) {
	c.address = address
}

func (c *customer) builtInAddressUpdate(address string) {
	c.address = address
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

type contact struct {
	firstname, lastname, phone string
}

type business struct {
	name, address string
	contact
}

func (b business) info() {
	pf("Contact for business %s is : %s %s\n", b.name, b.contact.firstname, b.contact.lastname)
}

func main() {

	// var myMap map [keyType]valueType

	var heroes map[string]string
	heroes = make(map[string]string)
	pl(heroes)

	villains := make(map[string]string)

	heroes["Batman"] = "B Wayne"
	heroes["Pigman"] = "Your Neighbour"
	heroes["The Flash"] = "B Allen"
	villains["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound", 3: "SomeOther"}
	pf("Batman is %v\n", heroes["Batman"])
	pf("Catperson is %v\n", heroes["Catperson"])
	pf("SuperPet 1 : %v\n", superPets[1])

	_, petExists := superPets[3]
	pl("Is there a third superPet : ", petExists)

	// structs

	var tS customer
	tS.name = "Tom Smith"
	tS.address = "4 Main St"
	tS.bal = 234.55

	getCustomerInfo(tS)
	updateCustAddress(&tS, "123 Nord Street, Town Central")
	getCustomerInfo(tS)
	tS.builtInAddressUpdate("666 Downtown Hell")
	getCustomerInfo(tS)

	sS := customer{"Sally Smith", "333 Death Road", 0.0}
	pl("Name of new customer :", sS.name)

	pl("-----------------------------")

	rect1 := rectangle{10.0, 14.3}
	pl("Rect Area : ", rect1.Area())

	pl("-----------------------------")

	con1 := contact{
		"James",
		"Wang",
		"5555-66666"}

	bus1 := business{
		"ABC Plumbing",
		"234 North Street",
		con1}

	bus1.info()

	pl("-----------------------------")

	var kitty Aniamal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	kitty2 = Cat("New Kitty")
	pl("Cats Name :", kitty2.Name())
	pl("kitty is of type :", typeOf(kitty))
	pl("kitty2 is of type :", typeOf(kitty2))
}
