package piscine

import "fmt"

func QuadA(x, y int) {
	// Return if x or y is not positive
	if x <= 0 || y <= 0 {
		return
	}

	// Print rectangle row by row
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top-left or bottom-left corner
			if (i == 0 || i == y-1) && j == 0 {
				fmt.Print("o")
				// Top-right or bottom-right corner
			} else if (i == 0 || i == y-1) && j == x-1 {
				fmt.Print("o")
				// Top or bottom edge
			} else if i == 0 || i == y-1 {
				fmt.Print("-")
				// Left or right edge
			} else if j == 0 || j == x-1 {
				fmt.Print("|")
				// Interior
			} else {
				fmt.Print(" ")
			}
		}
		// New line at the end of each row
		fmt.Println()
	}
}
