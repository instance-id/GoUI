package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/cmd"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CreateViewMenu() {

	FrameMenu = ui.CreateFrame(WindowMain, 25, 5, ui.BorderNone, ui.Fixed)
	FrameMenu.SetPack(ui.Vertical)

	FrameMain = ui.CreateFrame(FrameMenu, 5, 5, ui.BorderThin, ui.AutoSize)
	FrameMain.SetPack(ui.Vertical)
	FrameMain.SetTitle(TxtMainMenu)
	FrameMain.SetPaddings(2, 2)

	// --- Run Verifier --------------------------------------------------
	BtnRunVerifier = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtRunVerifier, ui.Fixed)
	BtnRunVerifier.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Main Settings -------------------------------------------------
	BtnMainSettings = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtMainSettings+TxtActive, ui.Fixed)
	BtnMainSettings.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Asset Settings ------------------------------------------------
	BtnDiscordSettings = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtDiscordSettings, ui.Fixed)
	BtnDiscordSettings.OnClick(func(ev ui.Event) {
		if !FrmDiscordSettings.Visible() {
			CommandDiscordSettings()
		}
	})

	// --- Plugins -------------------------------------------------------
	BtnPlugins = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtPlugins, ui.Fixed)
	BtnPlugins.OnClick(func(ev ui.Event) {
		if !FrmPlugins.Visible() {
			CommandPlugins()
		}
	})

	// --- Select Theme --------------------------------------------------
	BtnTheme = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtSelectTheme, ui.Fixed)
	BtnTheme.OnClick(func(ev ui.Event) {
		BtnTheme.SetEnabled(false)
		ChangeTheme(BtnTheme)
	})

	// --- Quit ----------------------------------------------------------
	BtnQuit = ui.CreateButton(FrameMain, ui.AutoSize, ui.AutoSize, TxtQuit, ui.Fixed)
	BtnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})
}
