package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

const st = tcell.StyleDefault

var char = [10]tcell.Style{
	st.Background(tcell.ColorBlack),
	st.Background(tcell.ColorMaroon),
	st.Background(tcell.ColorGreen),
	st.Background(tcell.ColorOlive),
	st.Background(tcell.ColorNavy),
	st.Background(tcell.ColorPurple),
	st.Background(tcell.ColorTeal),
	st.Background(tcell.ColorSilver),
	st.Background(tcell.ColorGray),
}

func main() {
	encoding.Register()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}

	if e = s.Init(); e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}

	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			event := s.PollEvent()
			switch event := event.(type) {
			case *tcell.EventKey:
				switch event.Key() {
				case tcell.KeyEscape:
					close(quit)
					return
				}
			case *tcell.EventResize:
				draw(s)
			}
		}
	}()

	draw(s)

	select {
	case <-quit:
		break
	}

	s.Fini()

}

func draw(s tcell.Screen) {
	w, h := s.Size()
	if w == 0 || h == 0 {
		return
	}

	data := genData(w, h)
	c := '0'
	for y := range data {
		for x := range data[y] {
			c = rune('0' + data[y][x]%10)
			s.SetCell(x, y, char[data[y][x]%10], c)
		}
	}

	s.Show()
}

func genData(w, h int) [][]int {
	data := make([][]int, h)
	for i := 0; i < h; i++ {
		data[i] = make([]int, w)
	}

	for y := range data {
		for x := range data[y] {
			if x > y {
				data[y][x] = x
			} else {
				data[y][x] = y
			}
		}
	}

	// fmt.Printf("\n%v\n", data)

	return data
}
