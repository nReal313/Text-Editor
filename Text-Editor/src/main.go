package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	//Initialize screen and handle errors
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen")
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Error initializing screen")
	}
	defer screen.Fini()

	screen.Clear() //clear screen

	//Display message
	putStr(screen, 0, 0, "Press ESC to exit")

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				return
			}
			putStr(screen, 0, 1, fmt.Sprintf("You pressed : %v", ev.Rune()))
			screen.Show()
		}
	}
}

//putStr function to put text on the screen

func putStr(s tcell.Screen, x, y int, str string) {
	for i, r := range str {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}

// logic to store the typed letters

//logic to delete the typed letters

//logic to navigate to a particular letter and delete or append

//logic to
