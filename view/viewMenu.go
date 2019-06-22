package view

import (
	"runtime"
	"sync"

	. "github.com/instance-id/GoUI/text"
	ui "github.com/instance-id/clui"
)

var wg sync.WaitGroup

func CreateViewMenu() {

	FrameMenu = ui.CreateFrame(WindowMain, 25, 5, ui.BorderNone, ui.Fixed)
	FrameMenu.SetPack(ui.Vertical)
	FrameMenu.SetBackColor(235)

	FrameMainMenu = ui.CreateFrame(FrameMenu, 5, 5, ui.BorderThin, ui.AutoSize)
	FrameMainMenu.SetPack(ui.Vertical)
	FrameMainMenu.SetTitle(TxtMainMenu)
	FrameMainMenu.SetPaddings(2, 2)
	FrameMainMenu.SetBackColor(235)

	// --- Run Verifier --------------------------------------------------
	BtnRunVerifier = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtRunVerifier, ui.Fixed)
	BtnRunVerifier.SetAlign(ui.AlignLeft)
	BtnRunVerifier.SetShadowType(ui.ShadowHalf)
	BtnRunVerifier.OnClick(func(ev ui.Event) {
		if !FrmVerifier.Visible() {
			CommandRunVerifier()
		}
	})

	// --- Main Settings -------------------------------------------------
	BtnMainSettings = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtMainSettings+TxtActive, ui.Fixed)
	BtnMainSettings.SetAlign(ui.AlignLeft)
	BtnMainSettings.SetShadowType(ui.ShadowHalf)
	BtnMainSettings.OnClick(func(ev ui.Event) {
		if !FrmMainSettings.Visible() {
			if LogRunning {
				LogRunning = false
			}
			CommandMainSettings()
		}
	})

	// --- Discord Settings ----------------------------------------------
	BtnDiscordSettings = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtDiscordSettings, ui.AutoSize)
	BtnDiscordSettings.SetAlign(ui.AlignLeft)
	BtnDiscordSettings.SetShadowType(ui.ShadowHalf)
	BtnDiscordSettings.OnClick(func(ev ui.Event) {
		if !FrmDiscordSettings.Visible() {
			if LogRunning {
				LogRunning = false
			}
			CommandDiscordSettings()
		}
	})

	// --- Database Settings ---------------------------------------------
	BtnDatabaseSettings = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtDatabaseSettings, ui.Fixed)
	BtnDatabaseSettings.SetAlign(ui.AlignLeft)
	BtnDatabaseSettings.SetShadowType(ui.ShadowHalf)
	BtnDatabaseSettings.OnClick(func(ev ui.Event) {
		if !FrmDatabaseSettings.Visible() {
			if LogRunning {
				LogRunning = false
			}
			CommandDatabaseSettings()
		}
	})

	// --- Plugins -------------------------------------------------------
	BtnPlugins = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtPlugins, ui.Fixed)
	BtnPlugins.SetAlign(ui.AlignLeft)
	BtnPlugins.SetShadowType(ui.ShadowHalf)
	BtnPlugins.OnClick(func(ev ui.Event) {
		if !FrmPlugins.Visible() {
			if LogRunning {
				LogRunning = false
			}
			CommandPlugins()
		}
	})

	// --- Logs ----------------------------------------------------------
	BtnLogs = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, Txtlogs, ui.Fixed)
	BtnLogs.SetAlign(ui.AlignLeft)
	BtnLogs.SetShadowType(ui.ShadowHalf)
	BtnLogs.OnClick(func(ev ui.Event) {
		if !LogRunningView {
			go LoadLogs()
			runtime.Gosched()
		}
		log := CreateLogDialog(Txtlogs)
		log.View.OnClose(func(ev ui.Event) bool {
			LogChan <- true
			return true
		})

	})

	// --- Quit ----------------------------------------------------------
	BtnQuit = ui.CreateButton(FrameMainMenu, 22, ui.AutoSize, TxtQuit, ui.Fixed)
	BtnQuit.SetAlign(ui.AlignLeft)
	BtnQuit.SetShadowType(ui.ShadowHalf)
	BtnQuit.OnClick(func(ev ui.Event) {
		if LogRunning {
			//LogChan <- true
			LogRunning = false
		}
		go ui.Stop()
	})
}
