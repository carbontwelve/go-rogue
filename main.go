package main

import "./bearlibterminal/BearLibTerminal"

func main() {
	glt.Open()
	defer glt.Close()
	glt.Print(0, 0, "Hello, world!")
	glt.Refresh()
	glt.Delay(1000)
}