package main

import (
	"fmt"
	"reflect"
)

type Farm struct {
	animals []Animal
}

func (f *Farm) Add(a ...Animal) {
	for _, x := range a {
		f.animals = append(f.animals, x)
	}
}

func (f *Farm) Describe() string {
	var (
		info         string = "The Farm consists of:\n"
		foodQuantity float64
	)
	for _, a := range f.animals {
		info += fmt.Sprintf("%v that weight %v kg and eats %v kg of food a month\n", a.typeName(), a.getWeight(), a.calcFood())
		foodQuantity += a.calcFood()
	}
	info += fmt.Sprintf("TOTAL food demand per month: %v kg\n", foodQuantity)
	return info
}

type (
	Animal interface {
		typeName() string
		getWeight() float64
		calcFood() float64
	}
)

type (
	Dog struct {
		weight float64
		mad    bool
	}

	Cat struct {
		weight float64
		fluffy bool
	}

	Cow struct {
		weight       float64
		milkQuantity int
	}
)

// GetWeight of animal
func (d *Dog) getWeight() float64 {
	return d.weight
}

func (c *Cat) getWeight() float64 {
	return c.weight
}

func (c *Cow) getWeight() float64 {
	return c.weight
}

//Calculate food for each animal
func (d *Dog) calcFood() float64 {
	const (
		voracity float64 = 10 / 5
		kMadness         = 10
	)
	if d.mad {
		return kMadness * d.weight * voracity
	}
	return d.weight * voracity
}

func (c *Cat) calcFood() float64 {
	const voracity float64 = 7
	return c.weight * voracity
}

func (c *Cow) calcFood() float64 {
	const voracity float64 = 25
	return c.weight * voracity
}

// GetType of animal for each species using black magic
func getType(i interface{}) string {
	value := reflect.ValueOf(i)
	typeName := reflect.Indirect(value).Type().Name()
	return fmt.Sprintf(typeName)
}

func (c *Cow) typeName() string {
	return getType(c)
}

func (d *Dog) typeName() string {
	return getType(d)
}

func (c *Cat) typeName() string {
	return getType(c)
}
