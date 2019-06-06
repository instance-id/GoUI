package elements

import (
	ui "github.com/VladimirMarkelov/clui"
)

// --- Application Elements ----------------------------------------------
var (
	WindowMain *ui.Window

	Btn1 *ui.Button
	Btn2 *ui.Button
	Btn3 *ui.Button

	FrameMenu          *ui.Frame
	FrmMainSettings    *ui.Frame
	FrmDiscordSettings *ui.Frame
	FrmPlugins         *ui.Frame

	// --- Menu buttons -------------------------
	BtnRunVerifier     *ui.Button
	BtnMainSettings    *ui.Button
	BtnDiscordSettings *ui.Button
	BtnPlugins         *ui.Button
	BtnTheme           *ui.Button
	BtnQuit            *ui.Button

	// --- Main Settings buttons ----------------
	BtnSave     *ui.Button
	BtnLogLevel *ui.Button
)
