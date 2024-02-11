package main

import "os"

func main() {
	os.Exit(1) // want "prohibited expression os.Exit"
}
