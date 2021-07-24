package main

import "github.com/01-edu/z01"

func main() {
	door := &Door{
		state: "OPEN",
	}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}

// Door struct
type Door struct {
	state string
}

// PrintStr function print a string using PrintRune function.
func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

// OpenDoor function
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	PrintStr("\n")
	ptrDoor.state = "OPEN"
}

// CloseDoor function
func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	PrintStr("\n")
	ptrDoor.state = "CLOSE"
}

// IsDoorOpen function
func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	PrintStr("\n")
	if ptrDoor.state == "OPEN" {
		return true
	}
	return false
}

// IsDoorClose function
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	PrintStr("\n")
	if ptrDoor.state == "CLOSE" {
		return true
	}
	return false
}