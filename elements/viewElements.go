package elements

import (
	ui "github.com/instance-id/clui"
)

// --- Application Elements ----------------------------------------------
var (
	WindowMain       *ui.Window
	WindowAssetCodes *ui.Window
	WindowAssetData  *ui.Window

	FrameMenu           *ui.Frame
	FrameMainMenu       *ui.Frame
	FrameContent        *ui.Frame
	FrmMainSettings     *ui.Frame
	FrmDiscordSettings  *ui.Frame
	FrmDatabaseSettings *ui.Frame
	FrmPlugins          *ui.Frame

	// --- Menu buttons -------------------------
	BtnRunVerifier      *ui.Button
	BtnMainSettings     *ui.Button
	BtnDiscordSettings  *ui.Button
	BtnDatabaseSettings *ui.Button
	BtnPlugins          *ui.Button
	BtnLogs             *ui.Button
	BtnQuit             *ui.Button

	// --- Main Settings buttons ----------------
	BtnMainSettingsSave *ui.Button
	BtnLogLevel         *ui.Button

	// --- Discord settings buttons -------------
	BtnAssetCodes   *ui.Button
	BtnAssetDetails *ui.Button

	// --- Database Settings buttons ----------------
	BtnDatabaseProvider *ui.Button

	// --- Arbitrary ------------------------------------------------
	WindowListDialog *ui.Window
	FrameListDialog  *ui.Frame
)
