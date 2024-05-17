package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}

type Zoo struct {
	Name    string
	Animals []*Animal
	Cages   []*Cage
}

type Animal struct {
	Icon         rune
	Species      string
	Name         string
	Age          int
	Size         string
	Caged        bool
	DamagePoints int
}

type Zookeper struct {
	Person
	XP int
}

type Cage struct {
	Size        string
	Species     string
	AnimalCount int
}

var animals = []*Animal{
	{Species: "lion", Name: "Alex", Age: 10, Size: "Big", Caged: false, DamagePoints: 30, Icon: '\U0001f981'},
	{Species: "zebra", Name: "Marty", Age: 10, Size: "Medium", Caged: false, DamagePoints: 20, Icon: '\U0001f993'},
	{Species: "giraffe", Name: "Melman", Age: 11, Size: "Big", Caged: false, DamagePoints: 10, Icon: '\U0001f992'},
	{Species: "hippopotamus", Name: "Gloria", Age: 10, Size: "Big", Caged: false, DamagePoints: 30, Icon: '\U0001f99b'},
	{Species: "chimpanzee", Name: "Mason", Age: 5, Size: "Small", Caged: false, DamagePoints: 15, Icon: '\U0001f412'},
	{Species: "chimpanzee", Name: "Phil", Age: 4, Size: "Small", Caged: false, DamagePoints: 15, Icon: '\U0001f412'},
}

var cages = []*Cage{
	{
		Size:        "big",
		Species:     "lion",
		AnimalCount: 0,
	},
	{
		Size:        "big",
		Species:     "zebra",
		AnimalCount: 0,
	},
	{
		Size:        "big",
		Species:     "giraffe",
		AnimalCount: 0,
	},
	{
		Size:        "big",
		Species:     "hippopotamus",
		AnimalCount: 0,
	},
	{
		Size:        "medium",
		Species:     "chimpanzee",
		AnimalCount: 0,
	},
}

func main() {
	zoo := Zoo{
		Name:  "Central Park Zoo",
		Cages: cages,
	}

	zookeeperDavid := Zookeper{
		Person: Person{
			Name:    "David",
			Surname: "Zoolander",
			Age:     25,
		},
		XP: 100,
	}

	zoo.greetVisitors()

	for true {
		if zookeeperDavid.XP <= 0 {
			fmt.Printf("Game over!")
			break
		}

		var escapedAnimals []*Animal

		for _, animal := range animals {
			if !animal.Caged {
				escapedAnimals = append(escapedAnimals, animal)
			}
		}

		if len(escapedAnimals) == 0 {
			fmt.Printf("All animals are now in cages! Good job, %s\n", zookeeperDavid.Name)
			break
		}

		randomAnimal := escapedAnimals[rand.Intn(len(escapedAnimals))]

		var cage *Cage

		for _, c := range zoo.Cages {
			if c.Species == randomAnimal.Species {
				cage = c
				break
			}
		}

		zookeeperDavid.catchAnimal(randomAnimal, cage)
	}
}

func (zookeper *Zookeper) catchAnimal(animal *Animal, cage *Cage) {
	randomInt := rand.Intn(2)

	switch randomInt {
	case 0:
		animal.biteZookeperAndScreamOutLoud(zookeper)
	case 1:
		fmt.Printf("%c Woah!! %s is now in cage\n", animal.Icon, animal.Name)

		cage.AnimalCount += 1
		animal.Caged = true

		if cage.AnimalCount == len(animals) {
			fmt.Printf("All animals are now in cage! Good job, %s\n", zookeper.Name)
		}
	}

}

func (zookeper *Zookeper) receiveDamageFromAnimal(animal *Animal) {
	zookeper.XP -= animal.DamagePoints

	if zookeper.XP > 0 {
		fmt.Printf("%s lost %v points of XP and now has %v points of XP\n", zookeper.Name, animal.DamagePoints, zookeper.XP)

	} else {
		fmt.Printf("%s lost %v points of XP and died\n", zookeper.Name, animal.DamagePoints)
	}
}

func (animal *Animal) biteZookeperAndScreamOutLoud(zookeper *Zookeper) {
	fmt.Printf("%c %s bites %s and screams out loud\n", animal.Icon, animal.Name, zookeper.Name)

	zookeper.receiveDamageFromAnimal(animal)
}

func (zoo *Zoo) greetVisitors() {
	fmt.Println("Hello! Welcome to the zoo. Unfortunately, today our animals decided to run away. Our zookeeper is trying to calm the situation down")
}
