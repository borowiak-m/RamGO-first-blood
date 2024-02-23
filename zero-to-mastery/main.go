package main

import (
	"fmt"
	"log"
)

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occypySpace(lot *ParkingLot, spaceNum int) {
	// change status to true for the passed space number
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	// change status to true for the passed space number
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

type Player struct {
	Health    byte
	MaxHealth byte
	Energy    byte
	MaxEnergy byte
	Name      string
}

func (p *Player) reportEnergy() {
	fmt.Println("Energy:", p.Energy)
	fmt.Println("Max Energy:", p.MaxEnergy)
	energy := float32(p.Energy)
	energymaximum := float32(p.MaxEnergy)
	energyUsage := energy / energymaximum * 100
	fmt.Printf("Energy used: %.2f\n", energyUsage)
	fmt.Println("Player", p.Name, "has", byte(energyUsage), "% of Energy used")
}

func (p *Player) reportHealth() {
	fmt.Println("Health:", p.Health)
	fmt.Println("Max Health:", p.MaxHealth)
	health := float32(p.Health)
	healthmaximum := float32(p.MaxHealth)
	healthUsage := health / healthmaximum * 100
	fmt.Printf("Energy used: %.2f\n", healthUsage)
	fmt.Println("Player", p.Name, "has", byte(healthUsage), "% of Health used")
}

func (p *Player) increaseEnergy(val byte) {
	p.Energy += val
	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}
	fmt.Println("Energy increased by", val)
	p.reportEnergy()
}
func (p *Player) decreaseEnergy(val byte) {
	if p.Energy-val > p.Energy {
		p.Energy = 0
	} else {
		p.Energy -= val
	}
	fmt.Println("Energy decreased by", val)
	p.reportEnergy()
}
func (p *Player) increaseHealth(val byte) {
	p.Health += val
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	fmt.Println("Health increased by", val)
	p.reportHealth()
}
func (p *Player) decreaseHealth(val byte) {
	if p.Health-val > p.Health {
		p.Health = 0
	} else {
		p.Health -= val
	}
	fmt.Println("Health decreased by", val)
	p.reportHealth()
}

func NewPlayer(health byte, maxHealth byte, energy byte, maxEnergy byte, name string) (*Player, error) {
	plr := &Player{
		Health:    health,
		MaxHealth: maxHealth,
		Energy:    energy,
		MaxEnergy: maxEnergy,
		Name:      name,
	}
	// error handle here
	return plr, nil
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("before\n", lot)
	lot.occupySpace(1)
	occypySpace(&lot, 2)
	fmt.Println("after occupy\n", lot)
	lot.vacateSpace(1)
	fmt.Println("after vacate\n", lot)
	fmt.Println("--------------")
	Adam, err := NewPlayer(150, 250, 120, 250, "Adam")
	if err != nil {
		log.Panic("SOMETHING WRONG WITH ADAM!!!")
	}
	fmt.Println(Adam)
	Adam.reportEnergy()
	Adam.reportHealth()
	Adam.increaseEnergy(50)
	Adam.decreaseHealth(100)
}
