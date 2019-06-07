package main

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	"github.com/instance-id/GoUI/view"
)

func InitData() {
	Log.LogLevel = []string{"INFO", "DEBUG", "WARNING", "ERROR"}
	Log.DefaultLogLevel = "INFO"
	Log.CurrentLogLevel = Log.DefaultLogLevel
	DiscordToken = "123123123SDFSDFSDFSDF1234123123"
	CommandPrefix = "!cmd "
}

func MainInitialSettings() {
	FrmMainSettings.SetVisible(true)
	FrmDiscordSettings.SetActive(false)
	FrmPlugins.SetActive(false)
}

func createView() {
	InitData()
	// --- Main Window ---------------------------------------------------
	WindowMain = ui.AddWindow(0, 0, 10, 7, TxtApplication)
	WindowMain.SetPack(ui.Horizontal)
	WindowMain.SetMaximized(true)

	// --- Main Menu Frame -----------------------------------------------
	view.CreateViewMenu()

	// --- Content Frame -------------------------------------------------
	view.CreateViewContent()

	// --- Settings Frames -----------------------------------------------
	view.CreateViewMainSettings()
	view.CreateViewDiscordSettings()
	view.CreateViewPlugins()

	// --- Popup Menu Frames ---------------------------------------------
	view.CreateViewPopupTheme()
	view.CreateViewLogLevel()
	//view.CreateViewPopupAssetCodes()

	MainInitialSettings()

	ui.MainLoop()
}

func mainLoop() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	ui.SetThemePath("../themes")

	createView()

	ui.MainLoop()
}

func main() {
	mainLoop()
}
