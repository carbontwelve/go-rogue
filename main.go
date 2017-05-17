package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

const animationSpeed = 10 * time.Millisecond
const backgroundColor = termbox.ColorBlue
const textColor = termbox.ColorWhite

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	render()

	// Game Loop
	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyEsc:
				return
			}
		}
		render()
		time.Sleep(animationSpeed)
	}
}

func render() {
	termbox.Clear(backgroundColor,backgroundColor)
	printText(0,0, textColor, backgroundColor, "Hello World!")
	termbox.Flush()
}

func printText(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}