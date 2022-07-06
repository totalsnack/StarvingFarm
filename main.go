package main

import "fmt"

func main() {
	var maxMad, mrKitty, filthyCow Animal

	maxMad = &Dog{
		weight: 14,
		mad:    true, //mad dog eats 10x more than regular one
	}

	mrKitty = &Cat{
		weight: 7,
		fluffy: true, //whatever
	}

	filthyCow = &Cow{
		weight:       200,
		milkQuantity: 200, // TODO: milk could be used for feeding cats
	}

	Farm := &Farm{}
	Farm.Add(maxMad, mrKitty, filthyCow)
	fmt.Println(Farm.Describe())
}
