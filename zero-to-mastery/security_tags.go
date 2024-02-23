package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main_security_tags() {
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	watch := Item{"Watch", Active}

	items := []Item{shirt, pants, purse, watch}
	fmt.Println("Initial\n", items)

	deactivate(&items[0].tag)
	fmt.Println("Item 0 deactivated\n", items)

	checkout(items)
	fmt.Println("Items checked out\n", items)
}
