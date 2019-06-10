package view

import (
	. "github.com/instance-id/GoUI/elements"
	ui "github.com/instance-id/clui"
)

func CreateViewContent() {
	FrameContent = ui.CreateFrame(WindowMain, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrameContent.SetPack(ui.Horizontal)
}
