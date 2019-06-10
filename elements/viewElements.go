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
	BtnRunVerifier      *ui.ButtonNoShadow
	BtnMainSettings     *ui.ButtonNoShadow
	BtnDiscordSettings  *ui.ButtonNoShadow
	BtnDatabaseSettings *ui.ButtonNoShadow
	BtnPlugins          *ui.ButtonNoShadow
	BtnTheme            *ui.ButtonNoShadow
	BtnLogs             *ui.ButtonNoShadow
	BtnQuit             *ui.ButtonNoShadow

	// --- Main Settings buttons ----------------
	BtnMainSettingsSave *ui.ButtonNoShadow
	BtnLogLevel         *ui.ButtonNoShadow

	// --- Discord settings buttons -------------
	BtnAssetCodes   *ui.ButtonNoShadow
	BtnAssetDetails *ui.ButtonNoShadow

	// --- Database Settings buttons ----------------
	BtnDatabaseProvider *ui.ButtonNoShadow

	// --- Arbitrary ------------------------------------------------
	WindowListDialog *ui.Window
	FrameListDialog  *ui.Frame
)
