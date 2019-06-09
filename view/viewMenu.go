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
	btnFrame := ui.CreateFrame(FrameMenu, ui.AutoSize, ui.AutoSize, ui.BorderThin, ui.Fixed)
	btnFrame.SetPaddings(1, 1)
	BtnRunVerifier = NewCreateButton(btnFrame, ui.AutoSize, ui.AutoSize, TxtRunVerifier, ui.Fixed)
	BtnRunVerifier.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Main Settings -------------------------------------------------
	BtnMainSettings = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtMainSettings+TxtActive, ui.Fixed)
	BtnMainSettings.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			CommandMainSettings()
		}
	})

	// --- Discord Settings ----------------------------------------------
	BtnDiscordSettings = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtDiscordSettings, ui.Fixed)
	BtnDiscordSettings.OnClick(func(ev ui.Event) {
		if !FrmDiscordSettings.Visible() {
			CommandDiscordSettings()
		}
	})

	// --- Database Settings ----------------------------------------------
	BtnDatabaseSettings = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtDatabaseSettings, ui.Fixed)
	BtnDatabaseSettings.OnClick(func(ev ui.Event) {
		if !FrmDatabaseSettings.Visible() {
			CommandDatabaseSettings()
		}
	})

	// --- Plugins -------------------------------------------------------
	BtnPlugins = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtPlugins, ui.Fixed)
	BtnPlugins.OnClick(func(ev ui.Event) {
		if !FrmPlugins.Visible() {
			CommandPlugins()
		}
	})

	// --- Select Theme --------------------------------------------------
	BtnTheme = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtSelectTheme, ui.Fixed)
	BtnTheme.OnClick(func(ev ui.Event) {
		BtnTheme.SetEnabled(false)
		ChangeTheme(BtnTheme)
	})

	// --- Quit ----------------------------------------------------------
	BtnQuit = ui.CreateButton(FrameMenu, ui.AutoSize, ui.AutoSize, TxtQuit, ui.Fixed)
	BtnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	FrameMenu.SetActive(true)
	FrameMain.SetActive(true)
	BtnRunVerifier.SetActive(true)
}
