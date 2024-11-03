package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var cur int = 0
var box []string

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
			} else if ev.Key() == tcell.KeyDelete || ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
				deleteStr()
			} else {
				storeStr(string(ev.Rune()))
			}
			showStr(screen)
		}
	}
}

//putStr function to put text on the screen

func putStr(s tcell.Screen, x, y int, str string) {
	for i, r := range str {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}

//logic to store the typed letters

func storeStr(str string) {
	box = append(box, str)
}

// logic to delete the typed letters
func deleteStr() {
	box = box[:len(box)-1]
}

//logic to navigate to a particular letter and delete or append

// logic to show the typed letters
func showStr(s tcell.Screen) {
	s.Clear()
	putStr(s, 0, 1, "Press ESC to exit, Backspace to delete.")
	var line string
	for _, r := range box {
		line += r
	}
	putStr(s, 0, 2, line)
	s.Show()
}
