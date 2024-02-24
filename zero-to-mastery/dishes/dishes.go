package main

import "fmt"

type Preperer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preperer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preperer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	prepareDishes(dishes)
}
