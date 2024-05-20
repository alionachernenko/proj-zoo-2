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
		Name:    "Central Park Zoo",
		Cages:   cages,
		Animals: animals,
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

	for {
		if zookeeperDavid.XP <= 0 {
			fmt.Printf("Game over!")
			break
		}

		var escapedAnimals []*Animal

		for _, animal := range zoo.Animals {
			if !animal.Caged {
				escapedAnimals = append(escapedAnimals, animal)
			}
		}

		if len(escapedAnimals) == 0 {
			fmt.Printf("All animals are now in cages! Good job, %s\n", zookeeperDavid.Name)
			break
		}

		randomAnimal := escapedAnimals[rand.Intn(len(escapedAnimals))]

		var currentCage *Cage

		for _, cage := range zoo.Cages {
			if cage.Species == randomAnimal.Species {
				currentCage = cage
				break
			}
		}

		zookeeperDavid.catchAnimal(randomAnimal, currentCage)
	}
}

func (zookeper *Zookeper) catchAnimal(animal *Animal, cage *Cage) {
	fmt.Printf("Look, It's %v %c. Catch it!\n", animal.Name, animal.Icon)

	randomInt := rand.Intn(2)

	var userInput int
	
	fmt.Scan(&userInput)

	if randomInt == userInput {
		fmt.Printf("%c Good job! %s is now in cage\n", animal.Icon, animal.Name)

		cage.AnimalCount += 1
		animal.Caged = true
	} else {
		animal.biteZookeperAndScreamOutLoud(zookeper)
	}
}

func (zookeper *Zookeper) receiveDamageFromAnimal(animal *Animal) {
	zookeper.XP -= animal.DamagePoints

	if zookeper.XP > 0 {
		fmt.Printf("%s loses %v points of XP and now has %v points of XP\n", zookeper.Name, animal.DamagePoints, zookeper.XP)

	} else {
		fmt.Printf("%s loses %v points of XP and dies :(\n", zookeper.Name, animal.DamagePoints)
	}
}

func (animal *Animal) biteZookeperAndScreamOutLoud(zookeper *Zookeper) {
	fmt.Printf("%c %s bites %s and screams out loud\n", animal.Icon, animal.Name, zookeper.Name)

	zookeper.receiveDamageFromAnimal(animal)
}

func (zoo *Zoo) greetVisitors() {
	fmt.Println("Hello! Welcome to the zoo. Unfortunately, today our animals decided to run away. Our zookeeper is trying to calm the situation down.\nRules: choose 1 or 0 to catch an animal")
}
