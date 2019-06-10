package utils

import (
	term "github.com/nsf/termbox-go"
)

type Screen struct {
	Buffer []term.Cell
	Width  int
	Height int
}

var screen Screen

func printLineAttr(x, y int, s string, fg, bg term.Attribute) {
	offsetX := 0
	offsetY := 0
	for _, char := range s {
		if char == '\n' {
			offsetX = 0
			offsetY++
		} else {
			term.SetCell(x+offsetX, y+offsetY, char, fg, bg)
			offsetX++
		}
	}
}

func printString(x, y int, s string) {
	// TODO срабатывает перенос на следующую строку при достижении правого края
	// и ошибка при выходе из границ буфера снизу
	offsetX := 0
	offsetY := 0
	for _, char := range s {
		if char == '\n' {
			offsetX = 0
			offsetY++
		} else {
			screen.Buffer[x+offsetX+(y+offsetY)*screen.Width].Ch = char
			offsetX++
		}
	}
}
