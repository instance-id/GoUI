package view

import (
	. "github.com/instance-id/GoUI/cmd"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	ui "github.com/instance-id/clui"
)

func CreateViewMenu() {

	FrameMenu = ui.CreateFrame(WindowMain, 25, 5, ui.BorderNone, ui.Fixed)
	FrameMenu.SetPack(ui.Vertical)

	FrameMainMenu = ui.CreateFrame(FrameMenu, 5, 5, ui.BorderThin, ui.AutoSize)
	FrameMainMenu.SetPack(ui.Vertical)
	FrameMainMenu.SetTitle(TxtMainMenu)
	FrameMainMenu.SetPaddings(2, 2)

	// --- Run Verifier --------------------------------------------------
	BtnRunVerifier = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtRunVerifier, ui.Fixed)
	BtnRunVerifier.SetAlign(ui.AlignLeft)
	BtnRunVerifier.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Main Settings -------------------------------------------------
	BtnMainSettings = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtMainSettings+TxtActive, ui.Fixed)
	BtnMainSettings.SetAlign(ui.AlignLeft)
	BtnMainSettings.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Discord Settings ----------------------------------------------
	BtnDiscordSettings = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtDiscordSettings, ui.AutoSize)
	BtnDiscordSettings.SetAlign(ui.AlignLeft)
	BtnDiscordSettings.OnClick(func(ev ui.Event) {
		if !FrmDiscordSettings.Visible() {
			CommandDiscordSettings()
		}
	})

	// --- Database Settings ----------------------------------------------
	BtnDatabaseSettings = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtDatabaseSettings, ui.Fixed)
	BtnDatabaseSettings.SetAlign(ui.AlignLeft)
	BtnDatabaseSettings.OnClick(func(ev ui.Event) {
		if !FrmDatabaseSettings.Visible() {
			CommandDatabaseSettings()
		}
	})

	// --- Plugins -------------------------------------------------------
	BtnPlugins = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtPlugins, ui.Fixed)
	BtnPlugins.SetAlign(ui.AlignLeft)
	BtnPlugins.OnClick(func(ev ui.Event) {
		if !FrmPlugins.Visible() {
			CommandPlugins()
		}
	})

	// --- Select Theme --------------------------------------------------
	BtnTheme = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtSelectTheme, ui.AutoSize)
	BtnTheme.SetAlign(ui.AlignLeft)
	BtnTheme.OnClick(func(ev ui.Event) {
		BtnTheme.SetEnabled(false)
		ChangeTheme(BtnTheme)
	})

	// --- Quit ----------------------------------------------------------
	BtnQuit = ui.CreateButton_NoShadow(FrameMainMenu, 22, ui.AutoSize, TxtQuit, ui.Fixed)
	BtnQuit.SetAlign(ui.AlignLeft)
	BtnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})
}
