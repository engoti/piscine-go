package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
	return true
}

func CloseDoor(ptrdoor *Door) bool {
	PrintStr("Door Closing...")
	ptrdoor.state = "OPEN"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Is the Door opened ?")
	ptrDoor.state = "CLOSE"
}

func IsDoorClose(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func main() {
	door := &Door{}
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
