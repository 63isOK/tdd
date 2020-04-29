package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var inputStr string = ""

func main() {
	fmt.Println("vim-go")

	encoding.Register()
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	sd := tcell.StyleDefault.Reverse(true)
	s.Clear()

	quit := make(chan struct{})
	newMessage := make(chan interface{})
	messages := make(chan string, 50)

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Clear()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					close(quit)
				} else if ev.Key() == tcell.KeyEnter {
					if len(inputStr) != 0 {
						newMessage <- 1
					}
				} else {
					inputStr += string(ev.Rune())
				}
			}
		}
	}()

	for {
		select {
		case <-quit:
			s.Fini()
			os.Exit(0)
		case <-time.After(time.Millisecond * 50):
		case <-newMessage:
			if len(inputStr) != 0 {
				messages <- inputStr
				inputStr = ""
			}
		}

		w, h := s.Size()
		drawBox(s, w, h, sd)

		s.Show()
	}
}

func drawBox(s tcell.Screen, w, h int, style tcell.Style) {
	if w <= 10 || h <= 10 {
		return
	}

	s.SetContent(0, 0, tcell.RuneULCorner, nil, style)
	s.SetContent(w-1, h-1, tcell.RuneLRCorner, nil, style)
	s.SetContent(0, h-1, tcell.RuneLLCorner, nil, style)
	s.SetContent(w-1, 0, tcell.RuneURCorner, nil, style)

	for i := 1; i < w-1; i++ {
		s.SetContent(i, 0, tcell.RuneHLine, nil, style)
		s.SetContent(i, h-1, tcell.RuneHLine, nil, style)
		s.SetContent(i, h-2, ' ', nil, tcell.StyleDefault)
		s.SetContent(i, h-3, tcell.RuneHLine, nil, style)
	}

	for i := 1; i < h-1; i++ {
		s.SetContent(0, i, tcell.RuneVLine, nil, style)
		s.SetContent(w-1, i, tcell.RuneVLine, nil, style)
	}

	s.SetContent(0, h-3, tcell.RuneLTee, nil, style)
	s.SetContent(w-1, h-3, tcell.RuneRTee, nil, style)
	drawString(s, 1, h-2, "intpu:")

	if len(inputStr) != 0 {
		drawString(s, 1+len("input: "), h-2, inputStr)
	}
}

func drawString(s tcell.Screen, w, h int, str string) {
	for index, char := range str {
		s.SetContent(w+index, h, char, nil, tcell.StyleDefault)
	}
}
