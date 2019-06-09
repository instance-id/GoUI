package elements

import (
	ui "github.com/VladimirMarkelov/clui"
)

// --- Application Elements ----------------------------------------------
var (
	WindowMain       *ui.Window
	WindowAssetCodes *ui.Window
	WindowAssetData  *ui.Window

	FrameMenu           *ui.Frame
	FrameMain           *ui.Frame
	FrameContent        *ui.Frame
	FrmMainSettings     *ui.Frame
	FrmDiscordSettings  *ui.Frame
	FrmDatabaseSettings *ui.Frame
	FrmPlugins          *ui.Frame

	// --- Menu buttons -------------------------
	BtnRunVerifier      *Button
	BtnMainSettings     *ui.Button
	BtnDiscordSettings  *ui.Button
	BtnDatabaseSettings *ui.Button
	BtnPlugins          *ui.Button
	BtnTheme            *ui.Button
	BtnQuit             *ui.Button

	// --- Main Settings buttons ----------------
	BtnMainSettingsSave *ui.Button
	BtnLogLevel         *ui.Button

	// --- Discord settings buttons -------------
	BtnAssetCodes *ui.Button

	// --- Arbitrary ------------------------------------------------
	WindowListDialog *ui.Window
	FrameListDialog  *ui.Frame
)
