package main

import "os"

func main() {
	for _, q := range os.Args {
		if q == "01" || q == "galaxy" || q == "galaxy 01" {
			os.Stdout.Write([]byte("Alert!!!\n"))
			os.Exit(0)
		}
	}
}
