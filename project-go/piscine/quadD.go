package piscine

import "fmt"

func QuadD(x, y int) {
	// Return if x or y is not positive
	if x <= 0 || y <= 0 {
		return
	}

	// Print rectangle row by row
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top-left corner
			if i == 0 && j == 0 {
				fmt.Print("A")
				// Top-right corner
			} else if i == 0 && j == x-1 {
				fmt.Print("C")
				// Bottom-left corner
			} else if i == y-1 && j == 0 {
				fmt.Print("A")
				// Bottom-right corner
			} else if i == y-1 && j == x-1 {
				fmt.Print("C")
				// Top or bottom edge
			} else if i == 0 || i == y-1 {
				fmt.Print("B")
				// Left or right edge
			} else if j == 0 || j == x-1 {
				fmt.Print("B")
				// Interior
			} else {
				fmt.Print(" ")
			}
		}
		// New line at the end of each row
		fmt.Println()
	}
}
