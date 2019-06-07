package utils

import (
	ui "github.com/VladimirMarkelov/clui"
)

type FramedInputParams struct {
	Orientation ui.PackType
	Height      int
	Width       int
	Border      ui.BorderStyle
	Scale       int
	PadX        int
	PadY        int
}

var (
	defaultOrientation = ui.PackType(ui.Vertical)
	defaultHeight      = ui.AutoSize
	defaultWidth       = ui.AutoSize
	defaultBorder      = ui.BorderThin
	defaultScale       = ui.Fixed
	defaultPadX        = 1
	defaultPadY        = 1
)

func NewFramedInput(parent *ui.Frame, title string, params *FramedInputParams) *ui.Frame {

	if params != nil {
		SetValues(params)
	}

	frameReturn := ui.CreateFrame(parent, defaultWidth, defaultHeight, defaultBorder, defaultScale)
	frameReturn.SetPaddings(defaultPadX, defaultPadY)
	frameReturn.SetPack(defaultOrientation)
	frameReturn.SetTitle(title)
	return frameReturn
}

func NewFramedWindowInput(parent *ui.Window, title string, params *FramedInputParams) *ui.Frame {

	if params != nil {
		SetValues(params)
	}

	frameWindowReturn := ui.CreateFrame(parent, defaultWidth, defaultHeight, defaultBorder, defaultScale)
	frameWindowReturn.SetPaddings(defaultPadX, defaultPadY)
	frameWindowReturn.SetPack(defaultOrientation)
	frameWindowReturn.SetTitle(title)
	return frameWindowReturn
}

func SetValues(params *FramedInputParams) {
	if params.Orientation != 1 {
		defaultOrientation = params.Orientation
	}
	if params.Height != 0 {
		defaultHeight = params.Height
	}
	if params.Width != 0 {
		defaultWidth = params.Width
	}
	if params.Border != 0 {
		defaultBorder = params.Border
	}
	if params.Scale != 0 {
		defaultScale = params.Scale
	}
	if params.PadX != 0 {
		defaultPadX = params.PadX
	}
	if params.PadY != 0 {
		defaultPadY = params.PadY
	}
}
