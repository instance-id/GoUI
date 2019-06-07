package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
)

func CreateViewContent() {
	FrameContent = ui.CreateFrame(WindowMain, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrameContent.SetPack(ui.Horizontal)
}
