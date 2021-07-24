package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	donnie := Pilot{}

	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = 1

	fmt.Println(donnie)
}
