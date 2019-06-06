package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/cmd"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CreateViewMenu() {

	FrameMenu = ui.CreateFrame(WindowMain, 16, 8, ui.BorderThin, ui.Fixed)
	FrameMenu.SetPack(ui.Vertical)
	FrameMenu.SetTitle(TxtMainMenu)
	FrameMenu.SetPaddings(2, 2)

	// --- Run Verifier --------------------------------------------------
	BtnRunVerifier = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtRunVerifier, 1)
	BtnRunVerifier.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Main Settings -------------------------------------------------
	BtnMainSettings = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtMainSettings+TxtActive, 1)
	BtnMainSettings.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Asset Settings ------------------------------------------------
	BtnDiscordSettings = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtDiscordSettings, 1)
	BtnDiscordSettings.OnClick(func(ev ui.Event) {
		if !FrmDiscordSettings.Visible() {
			CommandDiscordSettings()
		}
	})

	// --- Plugins -------------------------------------------------------
	BtnPlugins = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtPlugins, 1)
	BtnPlugins.OnClick(func(ev ui.Event) {
		if !FrmPlugins.Visible() {
			CommandPlugins()
		}
	})

	// --- Select Theme --------------------------------------------------
	BtnTheme = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtSelectTheme, 1)
	BtnTheme.OnClick(func(ev ui.Event) {
		BtnTheme.SetEnabled(false)
		ChangeTheme(BtnTheme)
	})

	// --- Quit ----------------------------------------------------------
	BtnQuit = ui.CreateButton(FrameMenu, ui.AutoSize, 3, TxtQuit, 1)
	BtnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})
}
