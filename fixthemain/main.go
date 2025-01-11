package main

import (
	"github.com/01-edu/z01"
)

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "closed"
}

func IsDoorOpen(ptrDoor *Door) string {
	PrintStr("is the Door opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) string {
	PrintStr("is the Door closed ?")
	return ptrDoor.state
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "opened"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) == "closed" {
		OpenDoor(door)
	}
	if IsDoorOpen(door) == "opened" {
		CloseDoor(door)
	}
	if door.state == "opened" {
		CloseDoor(door)
	}
}
